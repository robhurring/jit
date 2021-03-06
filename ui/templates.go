package ui

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/robhurring/jit/jit"
)

const (
	issueInfoTemplate = `
@y{{ .Key }}: @{!w}{{ .Fields.Summary }}
@{!k}{{ .URL }}

@bReporter:@| {{ if .Fields.Reporter }}{{ .Fields.Reporter.DisplayName }}{{end}}
@bAssigned:@| {{ if .Fields.Assignee }}{{ .Fields.Assignee.DisplayName }}{{end}}
@bDeveloper:@| {{ if .Fields.PrimaryDeveloper }}{{ .Fields.PrimaryDeveloper.DisplayName }}{{ end }}
@bReviewer:@| {{ if .Fields.CodeReviewer }}{{ .Fields.CodeReviewer.DisplayName }}{{ end }}
@bAssigned:@| {{ if .Fields.Assignee }}{{ .Fields.Assignee.DisplayName }}{{ end }}
{{ if .Links }}
@{!k}-----------------------8<-------------------------------------------------------@|

@yLinks ({{ len .Links }}):@|
{{ range $link := .Links }}
  @r{{ $link.Type }}@|
  {{ $link.Key }}: @{!k}[{{ $link.Status }}]@|:  @{!w}{{ $link.Summary | trim }}@|
{{ end }}
@{!k}-----------------------8<-------------------------------------------------------@|
{{ end }}
@{!m}Status:@| {{ .Fields.Status.Name }}

{{ .Fields.Description | trim }}
{{ if .Fields.Comment.Comments }}
@{!k}-----------------------8<-------------------------------------------------------@|

@yComments ({{ len .Fields.Comment.Comments }}):@|
{{ range $comment := .Fields.Comment.Comments }}
"{{ $comment.Body | trim }}"
@{!k}{{ $comment.Author.DisplayName }}@|
{{ end }}
{{ end }}`

	pullRequestTemplate = `
{{ if .CodeReviewer }}/cc {{ .CodeReviewer | username }}{{ end }}

[JIRA {{ .Key }}]({{ .URL }})
{{ if .Associated }}
### Associated
{{ range $associated := .Associated }}
{{ $associated }}{{ end }}
{{ end }}
### Summary

{{ if .ModifiedSpecs }}
### Modified Specs
{{ range $spec := .ModifiedSpecs }}
		bundle exec rspec {{ $spec }}{{ end }}
{{ else }}
### Testing

		bundle exec rspec
{{ end }}
{{ if .AddSignature }}




[![](https://raw.githubusercontent.com/robhurring/jit/master/badge.png)](https://github.com/robhurring/jit)
{{ end }}
`

	pullRequestInfoTemplate = `
@yThis pull request will merge @{Yk}{{ .Head }}@y into @{Yk}{{ .Base }}@|

@{!k}-----------------------8<-------------------------------------------------------@|

@{!w}{{ .Title }}@|

{{ .Body | escapeColors }}
`
)

var (
	templates *template.Template
)

func init() {
	templateFuncs := template.FuncMap{
		"trim":         strings.TrimSpace,
		"username":     jit.FindUsername,
		"escapeColors": escapeColors,
	}

	t := template.New("all")
	t, _ = t.New("issue.info").Funcs(templateFuncs).Parse(issueInfoTemplate)
	t, _ = t.New("pull-request.body").Funcs(templateFuncs).Parse(pullRequestTemplate)
	t, _ = t.New("pull-request.info").Funcs(templateFuncs).Parse(pullRequestInfoTemplate)

	templates = t
}

func RenderTemplate(name string, data interface{}) string {
	var b bytes.Buffer
	templates.ExecuteTemplate(&b, name, data)
	return b.String()
}

func PrintTemplate(name string, data interface{}) {
	Println(RenderTemplate(name, data))
}

func escapeColors(data string) string {
	return strings.Replace(data, "@", "@@", -1)
}
