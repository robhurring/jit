package commands

import (
	"github.com/codegangsta/cli"
	"github.com/github/hub/git"
	"github.com/robhurring/jit/utils"
)

func init() {
	CmdRunner.Add(&cli.Command{
		Name:      "branch",
		ShortName: "br",
		Usage:     "Create a new branch for the given ISSUE",
		Action: func(c *cli.Context) {
			// key, err := jit.FindIssueKey(c.Args())

			// if err != nil {
			// 	ui.Errorln(err)
			// } else {
			// issue := jit.GetIssue(key, false)
			// branchName := jit.IssueBranchName(issue)

			// ui.Println(branchName)
			utils.Debug(git.Head())
			// }
		},
	})
}
