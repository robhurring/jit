package git

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/robhurring/jit/cmd"
)

func Dir() (string, error) {
	output, err := cmd.New("git").WithArgs("rev-parse", "-q", "--git-dir").CombinedOutput()

	if err != nil {
		return "", fmt.Errorf("Not a git repository (or any of the parent directories): .git")
	}

	gitDir := strings.TrimSpace(string(output))
	gitDir, err = filepath.Abs(gitDir)
	if err != nil {
		return "", err
	}

	return gitDir, nil
}
