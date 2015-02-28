package main

import (
	"github.com/codegangsta/cli"
	"github.com/codeskyblue/go-sh"
)

type Commands struct {
}

func NewCommands() *Commands {
	c := &Commands{}
	return c
}

func (c *Commands) Info(cli *cli.Context) {
	args := cli.Args()

	if len(args) > 0 {
		issue := GetIssue(args[0], true)
		RenderTemplate("issue.info", issue)
	} else {
		// TODO: lookup issue from branch name
		Logger.Log("@rMissing ISSUE@|\nUsage: %s\n", cli.Command.Usage)
	}
}

func (c *Commands) Open(cli *cli.Context) {
	args := cli.Args()

	if len(args) > 0 {
		sh.Command("open", IssueURL(args[0])).Run()
	} else {
		// TODO: lookup issue from branch name
		Logger.Log("@rMissing ISSUE@|\nUsage: %s\n", cli.Command.Usage)
	}
}
