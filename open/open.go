package open

import (
	"github.com/codegangsta/cli"
	"github.com/codeskyblue/go-sh"
	"github.com/robhurring/jit/util"
)

var Commands []cli.Command

func init() {
	Commands = append(Commands, cli.Command{
		Name:      "open",
		ShortName: "o",
		Usage:     "open the given [ISSUE]",
		Action: func(c *cli.Context) {
			args := c.Args()

			if len(args) > 0 {
				Open(args[0])
			} else {
				// TODO: lookup issue from branch name
				util.Logger.Log("@rMissing ISSUE@|\nUsage: %s\n", c.Command.Usage)
			}
		},
	})
}

func Open(key string) {
	sh.Command("open", util.IssueURL(key)).Run()
}
