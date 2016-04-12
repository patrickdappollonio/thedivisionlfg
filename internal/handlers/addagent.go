package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/patrickdappollonio/thedivisionlfg/internal/models/player"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

func PostAddNew(_ context.Context, w http.ResponseWriter, r *http.Request) {
	// Create a new appengine ctx
	ctx := appengine.NewContext(r)

	// Create a dummy verification against spammers
	if r.Header.Get("X-Requested-With") != "XMLHttpRequest" {
		fmt.Fprintf(w, "")
		return
	}

	// Find the content of each variable
	var (
		username    = fnClean(r, "username")
		platform    = fnNum(r, "platform")
		activity    = fnNum(r, "activity")
		microphone  = fnNum(r, "microphone")
		lookingfor  = fnNum(r, "lookingfor")
		level       = fnNum(r, "level")
		dzlevel     = fnNum(r, "dzlevel")
		description = fnClean(r, "description")
	)

	// Create a player character
	current := player.Player{
		Username:     username,
		Platform:     player.Platform(platform),
		Activity:     player.Activity(activity),
		Microphone:   false,
		LookingFor:   player.LookingFor(lookingfor),
		StoryLevel:   level,
		DZLevel:      dzlevel,
		Description:  description,
		Signup:       time.Now(),
		IPAddress:    r.RemoteAddr,
		ForwardedFor: r.Header.Get("X-Forwarded-For"),
		DeletionID:   generateShortID(),
	}

	// Check if the agent has microphone
	if microphone == 1 {
		current.Microphone = true
	}

	// Verify the user and show error if corresponds
	if err := current.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a datastore key
	key := datastore.NewKey(ctx, player.DatastoreGroup, current.GetHash(), 0, nil)

	// Try finding first the user in the datastore
	var found player.Player
	if err := datastore.Get(ctx, key, &found); err != nil {
		if err != datastore.ErrNoSuchEntity {
			log.Errorf(ctx, err.Error())
			http.Error(w, "No se pudo registrar el agente en nuestro sistema. Intenta de nuevo más tarde.", http.StatusInternalServerError)
			return
		}
	}

	// Check if we got something
	if found.Username != "" {
		errmsg := fmt.Sprintf(`El agente "%s" ya está registrado en nuestra base de datos.`, found.Username)
		http.Error(w, errmsg, http.StatusConflict)
		return
	}

	// Store in datastore and check for error
	if _, err := datastore.Put(ctx, key, &current); err != nil {
		log.Errorf(ctx, err.Error())
		http.Error(w, "No se pudo registrar el agente en nuestro sistema. Intenta de nuevo más tarde.", http.StatusInternalServerError)
		return
	}

	// Print something if everything is okay
	response := KV{
		"ok":         true,
		"username":   current.Username,
		"deletionid": fmt.Sprintf(player.DeletionURLFormat, current.DeletionID),
	}

	// Create a placeholder
	content := new(bytes.Buffer)

	// Send the response back to the browser
	if err := json.NewEncoder(content).Encode(&response); err != nil {
		log.Errorf(ctx, err.Error())
		http.Error(w, "Un error desconocido ha ocurrido. Los administradores han sido notificados. Intenta agregar tu usuario más adelante.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, content.String())
}
