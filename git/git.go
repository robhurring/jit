package git

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/robhurring/jit/cmd"
)

func normalizeBranchName(branch string) (output string, ok bool) {
	ok, output = true, branch

	if branch == "" {
		ok = false
	}

	if strings.Contains(branch, "HEAD") {
		ok = false
	}

	if ok {
		pieces := strings.Split(branch, "/")
		output = pieces[len(pieces)-1]
	}

	return
}

func Dir() (string, error) {
	output, err := cmd.New("git").WithArgs("rev-parse", "-q", "--git-dir").CombinedOutput()

	if err != nil {
		return "", fmt.Errorf("Not a git repository (or any of the parent directories): .git")
	}

	gitDir := string(output[0])
	gitDir, err = filepath.Abs(gitDir)
	if err != nil {
		return "", err
	}

	return gitDir, nil
}

func BranchList() ([]string, error) {
	branchSet := make(map[string]int)

	branchCmd := cmd.New("git").WithArgs("branch", "-a", "--no-color")
	cutCmd := cmd.New("cut").WithArgs("-c", "3-")

	output, _, err := cmd.Pipeline(branchCmd, cutCmd)
	if err != nil {
		fmt.Println("erro.")
		return []string{}, err
	}

	list := strings.Split(string(output), "\n")
	for _, branch := range list {
		if normalized, ok := normalizeBranchName(branch); ok {
			if _, exists := branchSet[normalized]; !exists {
				branchSet[normalized] = 0
			}
		}
	}

	branches := []string{}
	for name := range branchSet {
		branches = append(branches, name)
	}

	return branches, err
}
