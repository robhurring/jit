package commands

import (
	"github.com/codegangsta/cli"
	"github.com/robhurring/jit/cmd"
	"github.com/robhurring/jit/jit"
)

func init() {
	CmdRunner.Add(&cli.Command{
		Name:      "open",
		ShortName: "o",
		Usage:     "Open the [ISSUE] in the browser",
		Action: func(c *cli.Context) {
			key, err := DetectIssue(c.Args())
			if err != nil {
				panic(err)
			}

			url := jit.IssueURL(key)
			err = cmd.Open(url)
			if err != nil {
				panic(err)
			}
		},
	})
}
