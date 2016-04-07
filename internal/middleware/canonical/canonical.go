package canonical

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pressly/chi"
	"golang.org/x/net/context"
)

func Enforce(hostname string, usessl bool) func(next chi.Handler) chi.Handler {
	return func(next chi.Handler) chi.Handler {
		fn := func(c context.Context, w http.ResponseWriter, r *http.Request) {
			if !strings.HasPrefix(r.Host, "localhost") {
				if r.Host != hostname {
					protocol := "http"

					if usessl {
						protocol = "https"
					}

					http.Redirect(w, r, fmt.Sprintf("%s://%s%s", protocol, hostname, r.URL.Path), http.StatusMovedPermanently)
					return
				}
			}

			next.ServeHTTPC(c, w, r)
		}

		return chi.HandlerFunc(fn)
	}
}
