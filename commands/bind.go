package commands

import (
	"strings"

	"github.com/codegangsta/cli"
	"github.com/robhurring/jit/git"
	"github.com/robhurring/jit/ui"
)

func init() {
	CmdRunner.Add(&cli.Command{
		Name:      "bind",
		ShortName: "b",
		Usage:     "Bind the current branch to [ISSUE]",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "list, l",
				Usage: "List issue binding for the current branch",
			},
		},
		Action: func(c *cli.Context) {
			branch, err := git.CurrentBranch()
			if err != nil {
				panic(err)
			}

			configPath := "branch." + branch.Name + ".issue"

			if c.Bool("l") {
				binding, err := git.GetConfig(configPath)
				if err != nil {
					panic(err)
				}

				ui.Printf("@{!w}Bound @{!k}->@| %s\n", strings.TrimSpace(binding))
				return
			}

			key, err := DetectIssue(c.Args())
			if err != nil {
				panic(err)
			}

			_, err = git.SetConfig(configPath, key)
			if err != nil {
				panic(err)
			}

			ui.Printf("@{!w}Bound @{!k}->@| %s\n", key)
		},
	})
}
