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
		"renderselect": func(max int) template.HTML {
			var elems []string
			for i := 1; i <= max; i++ {
				elems = append(elems, fmt.Sprintf("<option value=\"%v\">%v</option>\n", i, i))
			}
			return template.HTML(strings.Join(elems, ""))
		},
	},
}

func init() {
	Template = render.New(render.Options{
		Directory:     "views",
		Layout:        "layout",
		Extensions:    []string{".tmpl"},
		Funcs:         tmplfuncs,
		Charset:       "UTF-8",
		IsDevelopment: true,
	})
}
