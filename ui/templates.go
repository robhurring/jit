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
@{!k}{{ .Self }}

@bCreator:@|  {{ .Fields.Reporter.DisplayName }}
@bSponsor:@| {{ .Fields.Sponsor.DisplayName }}
@bDeveloper:@| {{ .Fields.PrimaryDeveloper.DisplayName }}
@bReviewer:@| {{ .Fields.CodeReviewer.DisplayName }}
@bQA Reviewer:@| {{ .Fields.QAReviewer.DisplayName }}
@bRelease Manager:@| {{ .Fields.ReleaseManager.DisplayName }}
@bAssigned:@| {{ .Fields.Assignee.DisplayName }}
{{ if .Links }}
@{!k}-----------------------8<-------------------------------------------------------@|

@yLinks ({{ len .Links }}):@|
{{ range $link := .Links }}
  @r{{ $link.Type }}@|
  {{ $link.Key }}: @{!k}[{{ $link.Status }}]@|:  @{!w}{{ $link.Summary | trim }}@|
{{ end }}
@{!k}-----------------------8<-------------------------------------------------------@|{{ end }}

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
/cc {{ .Fields.CodeReviewer.DisplayName | username }}

[JIRA {{ .Key }}]({{ .Self }}): {{ .Fields.Summary }}

### Associated

n/a

### Summary

* Changed A, B, C

### Testing

` + "`rake spec`"

	pullRequestInfoTemplate = `
@{!w}{{ .Title }}@|

{{ .Body }}
`
)

var (
	templates *template.Template
)

func init() {
	templateFuncs := template.FuncMap{
		"trim":     strings.TrimSpace,
		"username": findUsername,
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

func findUsername(name string) string {
	username := jit.FindUsername(name)
	return strings.Replace(username, "@", "@@", 1)
}
