package jit

import (
	"regexp"
	"strings"

	"github.com/robhurring/go-jira-client"
)

var (
	issueRe    = regexp.MustCompile("^([A-Z-]+-?[0-9]+)")
	jiraConfig *JiraConfig
)

func init() {
	jiraConfig = AppConfig.Jira
}

type Issue struct {
	*gojira.Issue
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

func ExtractIssue(data string) (match string, ok bool) {
	ok = false

	if found := issueRe.FindString(data); found != "" {
		ok = true
		match = found
	}

	return
}

func GetIssue(key string, allFields bool) (issue *Issue, err error) {
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

	jiraIssue, err := jira.Issue(key, params)
	if err != nil {
		return
	}

	issue = &Issue{
		Issue: jiraIssue,
	}

	return
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
