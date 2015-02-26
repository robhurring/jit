package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/robhurring/jit/branch"
	"github.com/robhurring/jit/info"
)

func main() {
	app := cli.NewApp()
	app.Name = "jit"
	app.Usage = "Jira + Git: A workflow story"
	app.Version = "0.0.1"
	app.Author = "Rob Hurring"
	app.Email = "robhurring@gmail.com"

	app.Commands = append(app.Commands, branch.Commands...)
	app.Commands = append(app.Commands, info.Commands...)

	app.Run(os.Args)
}
