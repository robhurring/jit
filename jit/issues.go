package jit

import (
	"strings"

	"github.com/robhurring/go-jira-client"
	"github.com/robhurring/jit/ui"
)

var jiraConfig *JiraConfig

func init() {
	jiraConfig = AppConfig.Jira
}

func FindIssueKey(args []string) (key string, err error) {
	err = nil
	key = ""

	if len(args) > 0 {
		key = NormalizeIssueKey(args[0])
	} else {
		// TODO: lookup issue from branch name
		err = ui.Error("No issue given, or could be found for the current branch!")
	}

	return
}

func GetIssue(key string, allFields bool) gojira.Issue {
	var params gojira.Params = nil
	key = NormalizeIssueKey(key)

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

	return jira.Issue(key, params)
}

func IssueURL(key string) string {
	key = NormalizeIssueKey(key)
	return jiraConfig.Host + "/browse/" + key
}

func NormalizeIssueKey(key string) string {
	if i := strings.Index(key, "-"); i == -1 {
		key = strings.ToUpper(jiraConfig.DefaultProject) + "-" + key
	}

	return key
}
