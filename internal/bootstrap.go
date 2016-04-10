package internal

import (
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"

	"github.com/patrickdappollonio/thedivisionlfg/internal/handlers"
)

var Router *chi.Mux

func init() {
	Router = chi.NewRouter()

	Router.Use(
		// canonical.Enforce(config.DomainName, config.ShouldUseSSL),
		middleware.Logger,
		// middleware.CloseNotify,
		middleware.Recoverer,
	)

	Router.Get("/", handlers.GetHome)
	Router.Post("/addnew", handlers.PostAddNew)
	Router.Post("/search", handlers.PostSearch)
}
