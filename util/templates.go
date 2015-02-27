package util

import (
	"bufio"
	"strings"
	"text/template"
	"time"
)

var (
	templates *template.Template
)

const issueInfoTemplate = `
@y{{ .Key }}: @{!w}{{ .Fields.Summary }}
@{!k}{{ .Self }}

@bCreator:@|  {{ .Fields.Reporter.DisplayName }}
@bSponsor:@| {{ .Fields.Sponsor.DisplayName }}
@bReviewer:@| {{ .Fields.CodeReviewer.DisplayName }}
@bAssigned:@| {{ .Fields.Assignee.DisplayName }}

@bStatus:@| {{ .Fields.Status.Name }}

{{ .Fields.Description | trim }}

{{ if .Fields.Comment.Comments }}
@yComments ({{ len .Fields.Comment.Comments }}):@|
{{ range $comment := .Fields.Comment.Comments }}
"{{ $comment.Body | trim }}"
@{!k}{{ $comment.Author.DisplayName }}@|
{{ end }}
{{ end }}`

func init() {
	templateFuncs := template.FuncMap{
		"trim": strings.TrimSpace,
		"time": func(t time.Time) string {
			return t.Format("Mon 3:04PM")
		},
	}

	t := template.New("all")
	t, _ = t.New("issue.info").Funcs(templateFuncs).Parse(issueInfoTemplate)

	templates = t
}

func RenderTemplate(name string, data interface{}) {
	// templates are rendered in chunks so we need to buffer it, otherwise
	// we lose the ability to use colors
	w := bufio.NewWriter(Logger)
	templates.ExecuteTemplate(w, name, data)
	w.Flush()
}
