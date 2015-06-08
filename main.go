package main

import (
	"github.com/codegangsta/cli"
	"github.com/robhurring/jit/commands"
)

const (
	Version = "1.1.0"
)

func main() {
	app := cli.NewApp()
	app.Name = "jit"
	app.Usage = "Jira + Git: A workflow story"
	app.Version = Version
	app.Author = "Rob Hurring"
	app.Email = "robhurring@gmail.com"
	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
	}
	commands.CmdRunner.Execute(app)
}
