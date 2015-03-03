package commands

import (
	"os"

	"github.com/codegangsta/cli"
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
	app.Commands = r.commands
	app.Run(os.Args)
}
