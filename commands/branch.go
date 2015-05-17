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

			key, err := DetectIssue(c.Args())
			if err != nil {
				panic(err)
			}

			issue, err := jit.GetIssue(key, true)
			if err != nil {
				panic(err)
			}

			branchName := issue.BranchName()

			if c.Bool("preview") {
				ui.Printf("@{!w}%s@|\n", branchName)
				return
			}

			if c.Bool("copy") {
				if err := cmd.Copy(branchName); err != nil {
					panic(err)
				}

				ui.Printf("@{!w}Copied!@| %s\n", branchName)
				return
			}

			branch := &git.Branch{Name: branchName}

			createBranch(branch)
		},
	})
}

func checkoutBranch(branch *git.Branch) {
	output, err := branch.Checkout()
	if err != nil {
		panic(err)
	}

	ui.Printf("@{!w}%s@|", output)
}

func createBranch(branch *git.Branch) {
	exists, err := branch.Exists()
	if err != nil {
		panic(err)
	}

	if exists {
		checkoutBranch(branch)
		return
	}

	if _, err := branch.Create(); err != nil {
		panic(err)
	}

	checkoutBranch(branch)
}
