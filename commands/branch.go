package commands

import (
	"github.com/codegangsta/cli"
	"github.com/robhurring/jit/cmd"
	"github.com/robhurring/jit/git"
	"github.com/robhurring/jit/jit"
	"github.com/robhurring/jit/ui"
)

func init() {
	CmdRunner.Add(&cli.Command{
		Name:      "branch",
		ShortName: "br",
		Usage:     "Create a new branch for the given ISSUE",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "preview, p",
				Usage: "Preview branch name",
			},
			cli.BoolFlag{
				Name:  "copy, c",
				Usage: "Copy branch name",
			},
		},
		Action: func(c *cli.Context) {
			// panic unless we have a .git dir
			if _, err := git.Dir(); err != nil {
				panic(err)
			}

			key, err := jit.FindIssueKey(c.Args())
			if err != nil {
				panic(err)
			}

			issue := jit.GetIssue(key, true)
			branch := jit.IssueBranchName(issue)

			if c.Bool("preview") {
				ui.Printf("@{!w}%s@|\n", branch)
				return
			}

			if c.Bool("copy") {
				if err := cmd.Copy(branch); err != nil {
					panic(err)
				}

				ui.Printf("@{!w}Copied!@| %s\n", branch)
				return
			}

			createBranch(branch)
		},
	})
}

func createBranch(branch string) {
	exists, err := git.BranchExists(branch)
	if err != nil {
		panic(err)
	}

	if exists {
		git.Checkout(branch)
	} else {
		_, err := git.CreateBranch(branch)
		if err != nil {
			panic(err)
		}

		git.Checkout(branch)
	}
}
