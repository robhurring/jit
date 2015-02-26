package info

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/robhurring/jit/util"
)

var Commands []cli.Command

func init() {
	Commands = append(Commands, cli.Command{
		Name:      "info",
		ShortName: "in",
		Usage:     "gather info about TICKET [or branch TICKET]",
		Action: func(c *cli.Context) {
			Info()
		},
	})
}

func Info() {
	config, _ := util.GetConfig()
	fmt.Print(config)
}
