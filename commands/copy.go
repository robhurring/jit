package commands

import (
	"github.com/codegangsta/cli"
	"github.com/codeskyblue/go-sh"
	"github.com/robhurring/jit/jit"
	"github.com/robhurring/jit/ui"
)

func init() {
	CmdRunner.Add(&cli.Command{
		Name:      "copy",
		ShortName: "cp",
		Usage:     "Copy the [ISSUE] URL to your clipboard",
		Action: func(c *cli.Context) {
			key, err := jit.FindIssueKey(c.Args())

			if err != nil {
				ui.Errorln(err)
			} else {
				url := jit.IssueURL(key)
				cmd := sh.Command("echo", url).Command("pbcopy")

				if err := cmd.Run(); err != nil {
					panic(err)
				} else {
					ui.Printf("@{!w}Copied!@| %s\n", url)
				}
			}
		},
	})
}
