package main

import (
	"net/http"

	"github.com/patrickdappollonio/thedivisionlfg/internal"
)

func init() {
	http.Handle("/", internal.Router)
}
