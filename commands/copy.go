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
				ui.Printf("@{!w}Copied!@| %s\n", url)
				sh.Command("echo", url).Command("pbcopy").Run()
			}
		},
	})
}
