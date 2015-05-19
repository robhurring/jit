package commands

import (
	"github.com/codegangsta/cli"
	"github.com/robhurring/go-jira-client"
	"github.com/robhurring/jit/jit"
	"github.com/robhurring/jit/ui"
)

func init() {
	CmdRunner.Add(&cli.Command{
		Name:      "info",
		ShortName: "in",
		Usage:     "Get info about an ISSUE",
		Action: func(c *cli.Context) {
			key, err := DetectIssue(c.Args())
			if err != nil {
				panic(err)
			}

			issue, err := jit.GetIssue(key, true)
			if err != nil {
				panic(err)
			}

			issueInfo(issue)
		},
	})
}

type infoData struct {
	*jit.Issue
	Links []*issueLink
	URL   string
}

type issueLink struct {
	Type    string
	Key     string
	Summary string
	Status  string
}

func issueInfo(issue *jit.Issue) {
	info := &infoData{Issue: issue}
	info.URL = issue.URL()

	links := make([]*issueLink, 0)

	for _, link := range issue.Fields.IssueLinks {
		var linkedIssue *gojira.Issue
		if link.InwardIssue != nil {
			linkedIssue = link.InwardIssue
		} else {
			linkedIssue = link.OutwardIssue
		}

		newLink := &issueLink{
			Type:    link.Type.Name,
			Key:     linkedIssue.Key,
			Summary: linkedIssue.Fields.Summary,
			Status:  linkedIssue.Fields.Status.Name,
		}

		links = append(links, newLink)
	}

	info.Links = links

	ui.PrintTemplate("issue.info", info)
}
