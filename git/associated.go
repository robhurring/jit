package git

import (
	"path"

	"github.com/robhurring/jit/jit"
	"github.com/robhurring/jit/utils"
)

func AssociatedProjects(match string) (associated []string) {
	config := jit.GetConfig()

	for _, associatedPath := range config.AssociatedPaths {
		utils.WalkTree(associatedPath, func(dir string) {
			if HasBranchNamed(dir, match) {
				associated = append(associated, path.Base(dir))
			}
		})
	}

	return
}
