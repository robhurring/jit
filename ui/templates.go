package ui

import (
	"bytes"
	"strings"
	"text/template"
)

const (
	issueInfoTemplate = `
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

	pullRequestTemplate = `
OHai.
`
)

var (
	templates *template.Template
)

func init() {
	templateFuncs := template.FuncMap{
		"trim": strings.TrimSpace,
	}

	t := template.New("all")
	t, _ = t.New("issue.info").Funcs(templateFuncs).Parse(issueInfoTemplate)

	templates = t
}

func RenderTemplate(name string, data interface{}) {
	var b bytes.Buffer
	templates.ExecuteTemplate(&b, name, data)
	Println(b.String())
}
