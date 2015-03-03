package commands

import "github.com/robhurring/jit/ui"

func FindIssueKey(args []string) (key string, err error) {
	err = nil
	key = ""

	if len(args) > 0 {
		key = args[0]
	} else {
		// TODO: lookup issue from branch name
		err = ui.Error("No issue given, or could be found for the current branch!")
	}

	return
}
