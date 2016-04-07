package handlers

import (
	"net/http"

	"github.com/patrickdappollonio/thedivisionlfg/internal/helpers/render"
	"golang.org/x/net/context"
)

func GetHome(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	render.Template.HTML(w, http.StatusOK, "home", nil)
}
