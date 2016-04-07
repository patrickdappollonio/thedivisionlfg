package render

import (
	"fmt"
	"html/template"
	"net/url"
	"strings"

	"github.com/unrolled/render"
)

var (
	Template *render.Render
)

var tmplfuncs = []template.FuncMap{
	template.FuncMap{
		"html": func(s string) template.HTML {
			return template.HTML(s)
		},
		"htmlattr": func(s string) template.HTMLAttr {
			return template.HTMLAttr(s)
		},
		"sprintf": func(f string, a ...interface{}) string {
			return fmt.Sprintf(f, a...)
		},
		"urlencode": func(s string) string {
			return url.QueryEscape(s)
		},
		"ifnotempty": func(value, def string) string {
			if strings.TrimSpace(value) != "" {
				return value
			}
			return def
		},
	},
}

func init() {
	Template = render.New(render.Options{
		Directory:     "views",
		Layout:        "layout",
		Extensions:    []string{".tmpl"},
		Charset:       "UTF-8",
		IsDevelopment: true,
	})
}
