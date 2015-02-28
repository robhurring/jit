package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	// defer util.SaveConfig()
	lazycmd := lazyCmd()

	app := cli.NewApp()
	app.Name = "jit"
	app.Usage = "Jira + Git: A workflow story"
	app.Version = "0.0.1"
	app.Author = "Rob Hurring"
	app.Email = "robhurring@gmail.com"
	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
	}
	app.Commands = *commands(lazycmd)

	app.Run(os.Args)
}

func commands(cmd lazyLoadCmd) *[]cli.Command {
	return &[]cli.Command{
		{
			Name:      "info",
			ShortName: "in",
			Usage:     "Get info about an ISSUE",
			Action:    func(c *cli.Context) { cmd().Info(c) },
		},
		{
			Name:      "open",
			ShortName: "o",
			Usage:     "Open an ISSUE in your browser",
			Action:    func(c *cli.Context) { cmd().Open(c) },
		},
	}
}

type lazyLoadCmd func() *Commands

func lazyCmd() lazyLoadCmd {
	var cmd *Commands
	return func() *Commands {
		if cmd == nil {
			cmd = NewCommands()
		}
		return cmd
	}
}
