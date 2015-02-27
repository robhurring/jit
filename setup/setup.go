package setup

import "github.com/robhurring/jit/util"

func Needed() (needed bool) {
	config := util.GetConfig()

	needed = false
	needed = !config.Jira.FilledOut()
	needed = !config.Github.FilledOut()

	return
}

func Run() {
	util.Logger.Log("@r[setup needed]\n\n")
}
