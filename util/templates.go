package util

import (
	"bufio"
	"strings"
	"text/template"
	"time"
)

var (
	templatesPath = "templates/*"
	templates     *template.Template
)

func init() {
	templateFuncs := template.FuncMap{
		"trim": strings.TrimSpace,
		"time": func(t time.Time) string {
			return t.Format("Mon 3:04PM")
		},
	}

	templates = template.Must(template.New("all").Funcs(templateFuncs).ParseGlob(templatesPath))
}

func RenderTemplate(name string, data interface{}) {
	// templates are rendered in chunks so we need to buffer it, otherwise
	// we lose the ability to use colors
	w := bufio.NewWriter(Logger)
	templates.ExecuteTemplate(w, name, data)
	w.Flush()
}
