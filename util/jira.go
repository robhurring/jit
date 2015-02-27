package util

import "github.com/robhurring/go-jira-client"

func GetIssue(id string) gojira.Issue {
	config := GetConfig()
	jiraConfig := config.Jira

	jira := gojira.NewJira(
		jiraConfig.Host,
		jiraConfig.ApiPath,
		jiraConfig.ActivityPath,
		&gojira.Auth{jiraConfig.Login, jiraConfig.Password},
	)

	return jira.Issue(id)
}
