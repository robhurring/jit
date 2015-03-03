package commands

import (
	"github.com/codegangsta/cli"
	"github.com/robhurring/jit/jit"
	"github.com/robhurring/jit/ui"
)

func init() {
	CmdRunner.Add(&cli.Command{
		Name:      "info",
		ShortName: "in",
		Usage:     "Get info about an ISSUE",
		Action: func(c *cli.Context) {
			key, err := FindIssueKey(c.Args())

			if err != nil {
				ui.Errorln(err)
			} else {
				issue := jit.GetIssue(key, true)
				ui.RenderTemplate("issue.info", issue)
			}
		},
	})
}
