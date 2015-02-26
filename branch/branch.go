package branch

import (
	"fmt"

	"github.com/codegangsta/cli"
)

var Commands []cli.Command

func init() {
	Commands = append(Commands, cli.Command{
		Name:      "branch",
		ShortName: "br",
		Usage:     "create a branch for the given TICKET",
		Action: func(c *cli.Context) {
			fmt.Println("Hello!")
		},
	})
}
