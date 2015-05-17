package jit

import (
	"regexp"
	"strings"

	"github.com/robhurring/go-jira-client"
)

var jiraConfig *JiraConfig

func init() {
	jiraConfig = AppConfig.Jira
}

type Issue struct {
	gojira.Issue
}

func (i *Issue) URL() string {
	return IssueURL(i.Key)
}

func (i *Issue) String() string {
	summary := i.Fields.Summary
	return strings.ToUpper(i.Key) + ": " + strings.ToLower(summary)
}

func (i *Issue) BranchName() string {
	maxBranchLength := AppConfig.MaxBranchLength
	summary := i.Fields.Summary

	re := regexp.MustCompile(`[^\w-]+`)
	fullName := strings.ToUpper(i.Key) + "_" + strings.ToLower(summary)
	underscored := re.ReplaceAllString(fullName, "_")

	return trimMaxLength(underscored, "_", maxBranchLength)
}

func FindIssueKey(args []string) (key string, err error) {
	err = nil
	key = ""

	if len(args) > 0 {
		key = NormalizeIssueKey(args[0])
	} else {
		panic("No issue given, or could be found for the current branch!")
	}

	return
}

func GetIssue(key string, allFields bool) *Issue {
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

	jiraIssue := jira.Issue(key, params)
	issue := &Issue{
		Issue: jiraIssue,
	}

	return issue
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

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func trimMaxLength(s, sep string, maxLength int) string {
	cutLine := maxLength

	if len(s) > maxLength {
		trimmed := s[:cutLine]
		lastOccurance := strings.LastIndex(trimmed, sep)
		cutLine = min(maxLength, lastOccurance)
		return trimmed[:cutLine]
	}

	return s
}
