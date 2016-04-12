package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/patrickdappollonio/thedivisionlfg/internal/helpers/render"
	"github.com/patrickdappollonio/thedivisionlfg/internal/models/player"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

func loadAgentByDeletionID(ctx context.Context, deletionID string) (*player.Player, *HTTPError) {
	// Create a query
	q := datastore.NewQuery(player.DatastoreGroup).Filter("DeletionID =", deletionID).Limit(1)

	// Perform the query
	var agent []player.Player
	if _, err := q.GetAll(ctx, &agent); err != nil {
		log.Errorf(ctx, err.Error())
		return nil, &HTTPError{http.StatusInternalServerError, fmt.Errorf("No se pudo encontrar al agente en el sistema. Ponte en contacto con el equipo en el Twitter: https://www.twitter.com/thedivisionlat")}
	}

	// Check if, by any weird rea
	if len(agent) == 0 {
		return nil, &HTTPError{http.StatusNotFound, fmt.Errorf("Código de eliminación no encontrado.")}
	}

	a := agent[0]
	return &a, nil
}

func RemoveAgent(c context.Context, w http.ResponseWriter, r *http.Request) {
	// Fetch the ID and check it
	deletionid, err := validateIDParam(c, "deletionid")

	// Check if the variable is right
	if err != nil {
		http.Error(w, err.Error.Error(), err.Status)
		return
	}

	// Create a new appengine ctx
	ctx := appengine.NewContext(r)

	// Load the agent
	agent, err := loadAgentByDeletionID(ctx, deletionid)

	// Check if it was possible
	if err != nil {
		http.Error(w, err.Error.Error(), err.Status)
		return
	}

	// Retrieve the first agent (should be the only one found)
	// and pass it to the template
	render.Template.HTML(w, http.StatusOK, "deleting", KV{
		"Agent": agent,
	})
}

func PostRemoveAgent(c context.Context, w http.ResponseWriter, r *http.Request) {
	// Fetch the ID and check it
	deletionid, err := validateIDParam(c, "deletionid")

	// Check if the variable is right
	if err != nil {
		http.Error(w, err.Error.Error(), err.Status)
		return
	}

	// Create a new appengine ctx
	ctx := appengine.NewContext(r)

	// Load the agent
	agent, err := loadAgentByDeletionID(ctx, deletionid)

	// Check if it was possible
	if err != nil {
		http.Error(w, err.Error.Error(), err.Status)
		return
	}

	// Check if the entered player is equal to the player we're looking for
	if strings.ToLower(fnClean(r, "username")) != strings.ToLower(agent.Username) {
		http.Error(w, "Lo siento, pero el nombre del Agente no corresponde con el nombre que tenemos guardado.", http.StatusNotFound)
		return
	}

	// Create a datastore key
	key := datastore.NewKey(ctx, player.DatastoreGroup, agent.GetHash(), 0, nil)

	// Try removing it
	if err := datastore.Delete(ctx, key); err != nil {
		log.Errorf(ctx, err.Error())
		http.Error(w, "No pudimos eliminar al agente de nuestra base de datos. Intenta más tarde.", http.StatusInternalServerError)
		return
	}

	render.Template.HTML(w, http.StatusOK, "deleted", KV{
		"Username": agent.Username,
	})
}
