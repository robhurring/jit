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
		Usage:     "gather info about ISSUE [or branch ISSUE]",
		Action: func(c *cli.Context) {
			args := c.Args()

			if len(args) > 0 {
				Info(args[0])
			} else {
				// TODO: lookup issue from branch name
				util.Logger.Log("@rMissing ISSUE@|\nUsage: %s\n", c.Command.Usage)
			}
		},
	})
}

func Info(key string) {
	issue := util.GetIssue(key, true)
	util.RenderTemplate("issue.info", issue)
}
