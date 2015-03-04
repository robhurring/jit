package commands

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/github/hub/ui"
)

var (
	CmdRunner = &Runner{}
)

type Runner struct {
	commands []cli.Command
}

func (r *Runner) Add(c *cli.Command) {
	r.commands = append(r.commands, *c)
}

func (r *Runner) Execute(app *cli.App) {
	defer func() {
		if err := recover(); err != nil {
			ui.Errorln(err)
			os.Exit(1)
		}
	}()

	app.Commands = r.commands
	app.Run(os.Args)
}
