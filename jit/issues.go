package jit

import (
	"regexp"
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
		key = jiraConfig.DefaultProject + "-" + key
	}

	return strings.ToUpper(key)
}

func IssueBranchName(issue gojira.Issue) string {
	maxBranchLength := AppConfig.MaxBranchLength

	re := regexp.MustCompile(`[\s]+`)
	fullName := strings.ToUpper(issue.Key) + "_" + strings.ToLower(issue.Fields.Summary)
	underscored := re.ReplaceAllString(fullName, "_")

	return trimMaxLength(underscored, "_", maxBranchLength)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func trimMaxLength(s, sep string, maxLength int) string {
	cutLine := maxLength

	trimmed := s[:cutLine]
	lastOccurance := strings.LastIndex(trimmed, sep)
	cutLine = min(maxLength, lastOccurance)

	return trimmed[:cutLine]
}
