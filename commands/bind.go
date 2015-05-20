package commands

import (
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
			cli.BoolFlag{
				Name:  "delete, d",
				Usage: "Delete the issue binding for the current branch",
			},
		},
		Action: func(c *cli.Context) {
			branch, err := git.CurrentBranch()
			if err != nil {
				panic(err)
			}

			configPath := "branch." + branch.Name + ".issue"

			if c.Bool("l") {
				binding, _ := git.GetConfig(configPath)
				if binding == "" {
					ui.Printf("@{!r}No bindings found.\n")
				} else {
					ui.Printf("@{!w}Bound @{!k}->@| %s\n", binding)
				}

				return
			}

			if c.Bool("d") {
				currentBinding, _ := git.GetConfig(configPath)

				if currentBinding == "" {
					ui.Printf("@{!r}No bindings found.\n")
				} else {
					git.UnsetConfig(configPath)
					ui.Printf("@{!r}Un-bound @{!k}->@| %s\n", currentBinding)
				}

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
