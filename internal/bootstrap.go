package internal

import (
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"

	"github.com/patrickdappollonio/thedivisionlfg/internal/config"
	"github.com/patrickdappollonio/thedivisionlfg/internal/handlers"
	"github.com/patrickdappollonio/thedivisionlfg/internal/middleware/canonical"
)

var Router *chi.Mux

func init() {
	Router = chi.NewRouter()

	Router.Use(
		canonical.Enforce(config.DomainName, config.ShouldUseSSL),
		middleware.Logger,
		middleware.CloseNotify,
		middleware.Recoverer,
	)

	Router.Get("/", handlers.GetHome)
}
