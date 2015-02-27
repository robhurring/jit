package info

import (
	"github.com/codegangsta/cli"
	"github.com/robhurring/jit/util"
)

var Commands []cli.Command

func init() {
	Commands = append(Commands, cli.Command{
		Name:      "info",
		ShortName: "in",
		Usage:     "gather info about TICKET [or branch TICKET]",
		Action: func(c *cli.Context) {
			Info()
		},
	})
}

func Info() {
	issue := util.GetIssue("AUTO-83")
	// util.Debug(issue.Fields.Comment.Comments[0].Created)
	util.RenderTemplate("issue.info", issue)
}
