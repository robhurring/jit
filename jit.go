package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/robhurring/jit/branch"
	"github.com/robhurring/jit/info"
	"github.com/robhurring/jit/open"
	"github.com/robhurring/jit/setup"
	"github.com/robhurring/jit/util"
)

func main() {
	defer util.SaveConfig()

	app := cli.NewApp()
	app.Name = "jit"
	app.Usage = "Jira + Git: A workflow story"
	app.Version = "0.0.1"
	app.Author = "Rob Hurring"
	app.Email = "robhurring@gmail.com"

	app.Commands = append(app.Commands, branch.Commands...)
	app.Commands = append(app.Commands, info.Commands...)
	app.Commands = append(app.Commands, open.Commands...)

	// Check if setup
	if setup.Needed() {
		setup.Run()
	}

	app.Run(os.Args)
}
