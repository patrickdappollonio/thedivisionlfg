package handlers

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/SKAhack/go-shortid"
	"github.com/pressly/chi"
	"golang.org/x/net/context"
)

type KV map[string]interface{}

const All = "all"

var reShortID = regexp.MustCompile(`[0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-]{10}`)

type HTTPError struct {
	Status int
	Error  error
}

func validateIDParam(c context.Context, param string) (string, *HTTPError) {
	// Find the deletion ID
	deletionID := chi.URLParam(c, param)

	// Validate if it's a proper deletion ID
	if utf8.RuneCountInString(deletionID) != 10 || !reShortID.MatchString(deletionID) {
		return "", &HTTPError{http.StatusNotFound, fmt.Errorf("Código de eliminación no encontrado.")}
	}

	return deletionID, nil
}

func fnClean(r *http.Request, param string) string {
	return strings.TrimSpace(r.FormValue(param))
}

func fnNum(r *http.Request, param string) int {
	clean := fnClean(r, param)

	if clean == All {
		return 0
	}

	n, _ := strconv.Atoi(clean)
	return n
}

func generateShortID() string {
	return shortid.Generator().Generate()
}
