package handlers

import (
	"fmt"
	"net/http"

	"golang.org/x/net/context"
)

func PostAddNew(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func PostSearch(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}
