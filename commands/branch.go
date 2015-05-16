package commands

import (
	"github.com/codegangsta/cli"
	"github.com/robhurring/jit/git"
	"github.com/robhurring/jit/jit"
	"github.com/robhurring/jit/utils"
)

func init() {
	CmdRunner.Add(&cli.Command{
		Name:      "branch",
		ShortName: "br",
		Usage:     "Create a new branch for the given ISSUE",
		Action: func(c *cli.Context) {
			// panic unless we have a .git dir
			if _, err := git.Dir(); err != nil {
				panic(err)
			}

			key, err := jit.FindIssueKey(c.Args())
			if err != nil {
				panic(err)
			}

			createBranch(key)
		},
	})
}

func createBranch(key string) {
	branches, err := git.BranchList()

	if err != nil {
		panic(err)
	}

	utils.Debug(branches)

}
