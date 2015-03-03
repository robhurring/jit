package jit

import (
	"strings"

	"github.com/robhurring/go-jira-client"
)

var jiraConfig *JiraConfig

func init() {
	jiraConfig = AppConfig.Jira
}

func GetIssue(key string, allFields bool) gojira.Issue {
	var params gojira.Params = nil

	jira := gojira.NewJira(
		jiraConfig.Host,
		jiraConfig.ApiPath,
		jiraConfig.ActivityPath,
		&gojira.Auth{jiraConfig.Login, jiraConfig.Password},
	)

	if !allFields {
		// basic fields list
		params = gojira.Params{"fields": "key,summary"}
	}

	return jira.Issue(NormalizeIssueKey(key), params)
}

func IssueURL(key string) string {
	return jiraConfig.Host + "/browse/" + NormalizeIssueKey(key)
}

func NormalizeIssueKey(key string) string {
	if i := strings.Index(key, "-"); i == -1 {
		key = jiraConfig.DefaultProject + "-" + key
	}

	return key
}
