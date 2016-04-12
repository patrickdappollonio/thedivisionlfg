package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"unicode/utf8"

	"github.com/patrickdappollonio/thedivisionlfg/internal/models/player"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

func PostSearch(_ context.Context, w http.ResponseWriter, r *http.Request) {
	// Create a new appengine ctx
	ctx := appengine.NewContext(r)

	// Retrieve all parameters
	var (
		platform   = fnNum(r, "platform")
		activity   = fnNum(r, "activity")
		microphone = fnNum(r, "microphone")
		lookingfor = fnNum(r, "lookingfor")
		level      = fnNum(r, "level")
		dzlevel    = fnNum(r, "dzlevel")
		search     = fnClean(r, "search")
	)

	// Convert all needed parameters to their correspondent elements
	// in Go
	var (
		mplatform   = player.Platform(platform)
		mactivity   = player.Activity(activity)
		mlookingfor = player.LookingFor(lookingfor)
		mlevel      = player.StoryLevel(level)
		mdzlevel    = player.DZLevel(dzlevel)
		mmicrophone = false
	)

	// Convert to the proper search value
	if microphone == 1 {
		mmicrophone = true
	} else if microphone == 2 {
		mmicrophone = false
	}

	// Check if platform exists
	if platform != 0 && mplatform.String() == "" {
		http.Error(w, "La plataforma que has seleccionado no existe", http.StatusBadRequest)
		return
	}

	// Check if activity exists
	if activity != 0 && mactivity.String() == "" {
		http.Error(w, "La actividad que has seleccionado no existe", http.StatusBadRequest)
		return
	}

	// Check if activity exists
	if lookingfor != 0 && mlookingfor.String() == "" {
		http.Error(w, "No has indicado una opción correcta sobre si estás buscando grupo o amistad", http.StatusBadRequest)
		return
	}

	// Check if activity exists
	if level != 0 && mlevel.String() == "" {
		http.Error(w, "No has indicado un tramo correcto para tu nivel en el modo historia", http.StatusBadRequest)
		return
	}

	// Check if activity exists
	if dzlevel != 0 && mdzlevel.String() == "" {
		http.Error(w, "No has indicado un tramo correcto para tu nivel en la Zona Oscura", http.StatusBadRequest)
		return
	}

	// Check if search is longer than 100 characters
	if utf8.RuneCountInString(search) > 100 {
		http.Error(w, "El tamaño de la cadena de texto excede el máximo permitido", http.StatusBadRequest)
		return
	}

	// Perform the search query
	q := datastore.NewQuery(player.DatastoreGroup)

	// If there's a platform selected
	if mplatform.String() != "" {
		q = q.Filter("Platform =", mplatform)
	}

	// If there's an activity selected
	if mactivity.String() != "" {
		q = q.Filter("Activity =", mactivity)
	}

	// If there's a LookingFor selected
	if mlookingfor.String() != "" {
		q = q.Filter("LookingFor =", mlookingfor)
	}

	// If there's a level selected
	if mlevel.String() != "" {
		if mlevel == player.SL_FIRST_TIER {
			q = q.Filter("StoryLevel <=", 10)
		} else if mlevel == player.SL_SECOND_TIER {
			q = q.Filter("StoryLevel >", 10).Filter("StoryLevel <=", 20)
		} else if mlevel == player.SL_THIRD_TIER {
			q = q.Filter("StoryLevel >", 20)
		}
	}

	// If there's a dzlevel selected
	if mdzlevel.String() != "" {
		if mdzlevel == player.DZ_FIRST_TIER {
			q = q.Filter("DZLevel <=", 30)
		} else if mdzlevel == player.DZ_SECOND_TIER {
			q = q.Filter("DZLevel >", 30).Filter("DZLevel <=", 50)
		} else if mdzlevel == player.DZ_THIRD_TIER {
			q = q.Filter("DZLevel >", 50)
		}
	}

	// If there's a microphone selected
	if microphone != 0 {
		log.Debugf(ctx, "Microphone !=0 ; Buscando por micrófono: %v", mmicrophone)
		q = q.Filter("Microphohe =", mmicrophone)
	}

	// Create a holder for agents
	var agents []player.Player

	// Perform the search query
	if _, err := q.GetAll(ctx, &agents); err != nil {
		log.Errorf(ctx, err.Error())
		http.Error(w, "No pudimos procesar la búsqueda. Intenta de nuevo más tarde.", http.StatusInternalServerError)
		return
	}

	// Create a placeholder
	content := new(bytes.Buffer)

	// Create a response
	response := KV{
		"agents": agents,
		"keywords": map[string][]string{
			"platform": []string{
				player.PC_PLATFORM.String(),
				player.PS4_PLATFORM.String(),
				player.XBOX_PLATFORM.String(),
			},
			"activity": {
				player.DAILY_CHALLENGING.String(),
				player.DAILY_HARD.String(),
				player.DARKZONE_FARMING.String(),
				player.DARKZONE_ROGUE.String(),
				player.STORY_NORMAL.String(),
				player.STORY_HARD.String(),
				player.STORY_CHALLENGE.String(),
				player.FREE_FARMING.String(),
				player.FREE_FREE.String(),
				player.FREE_SIDEMISSIONS.String(),
				player.TRADING.String(),
			},
			"lookingfor": {
				player.LOOKING_GROUP.String(),
				player.LOOKING_FRIENDS.String(),
			},
		},
	}

	// Send the response back to the browser
	if err := json.NewEncoder(content).Encode(&response); err != nil {
		log.Errorf(ctx, err.Error())
		http.Error(w, "No pudimos procesar la búsqueda. Intenta de nuevo más tarde.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, content.String())
}
