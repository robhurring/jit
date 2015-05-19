package commands

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/robhurring/jit/git"
	"github.com/robhurring/jit/jit"
	"github.com/robhurring/jit/ui"
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
	// Handle any panics
	defer func() {
		if err := recover(); err != nil {
			ui.Errorln(err)
			os.Exit(1)
		}
	}()

	// Update config
	defer func() {
		jit.SaveConfig()
	}()

	app.Commands = r.commands
	app.Run(os.Args)
}

func DetectIssue(args []string) (key string, err error) {
	if len(args) > 0 {
		key = jit.NormalizeIssueKey(args[0])
		return
	}

	branch, branchErr := git.CurrentBranch()
	if branchErr != nil {
		err = branchErr
		return
	}

	if match, ok := jit.ExtractIssue(branch.Name); ok {
		key = match
		return
	}

	keyPath := "branch." + branch.Name + ".issue"
	bound, err := git.GetConfig(keyPath)
	if err == nil {
		key = bound
		return
	}

	panic("No issue given, or found for the current branch!")
	return
}
