package commands

import (
	"github.com/codegangsta/cli"
	"github.com/codeskyblue/go-sh"
	"github.com/robhurring/jit/jit"
	"github.com/robhurring/jit/ui"
)

func init() {
	CmdRunner.Add(&cli.Command{
		Name:      "open",
		ShortName: "o",
		Usage:     "Open the [ISSUE] in the browser",
		Action: func(c *cli.Context) {
			key, err := jit.FindIssueKey(c.Args())

			if err != nil {
				ui.Errorln(err)
			} else {
				url := jit.IssueURL(key)
				cmd := sh.Command("open", url)

				if err := cmd.Run(); err != nil {
					panic(err)
				} else {
					ui.Printf("@{!w}Opening@| %s\n", url)
				}
			}
		},
	})
}
