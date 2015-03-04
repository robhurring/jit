package git

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/codeskyblue/go-sh"
)

func Dir() (string, error) {
	output, err := sh.Command("rev-parse", "-q", "--git-dir").Output()
	if err != nil {
		return "", fmt.Errorf("Not a git repository (or any of the parent directories): .git")
	}

	gitDir := output[0]
	gitDir, err = filepath.Abs(gitDir)
	if err != nil {
		return "", err
	}

	return gitDir, nil
}

func HasFile(segments ...string) bool {
	dir, err := Dir()
	if err != nil {
		return false
	}

	s := []string{dir}
	s = append(s, segments...)
	path := filepath.Join(s...)
	if _, err := os.Stat(path); err == nil {
		return true
	}

	return false
}

func BranchAtRef(paths ...string) (name string, err error) {
	dir, err := Dir()
	if err != nil {
		return
	}

	segments := []string{dir}
	segments = append(segments, paths...)
	path := filepath.Join(segments...)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	n := string(b)
	refPrefix := "ref: "
	if strings.HasPrefix(n, refPrefix) {
		name = strings.TrimPrefix(n, refPrefix)
		name = strings.TrimSpace(name)
	} else {
		err = fmt.Errorf("No branch info in %s: %s", path, n)
	}

	return
}

func Editor() (string, error) {
	output, err := sh.Command("var", "GIT_EDITOR").Output()
	if err != nil {
		return "", fmt.Errorf("Can't load git var: GIT_EDITOR")
	}

	return output[0], nil
}

func Head() (string, error) {
	return BranchAtRef("HEAD")
}

func SymbolicFullName(name string) (string, error) {
	output, err := sh.Command("rev-parse", "--symbolic-full-name", name).Output()
	if err != nil {
		return "", fmt.Errorf("Unknown revision or path not in the working tree: %s", name)
	}

	return output[0], nil
}

func Ref(ref string) (string, error) {
	output, err := sh.Command("rev-parse", "-q", ref).Output()
	if err != nil {
		return "", fmt.Errorf("Unknown revision or path not in the working tree: %s", ref)
	}

	return output[0], nil
}

func RefList(a, b string) ([]string, error) {
	ref := fmt.Sprintf("%s...%s", a, b)
	output, err := sh.Command("rev-list", "--cherry-pick", "--right-only", "--no-merges", ref).Output()
	if err != nil {
		return []string{}, fmt.Errorf("Can't load rev-list for %s", ref)
	}

	return output, nil
}

func Remotes() ([]string, error) {
	return sh.Command("remote", "-v").Output()
}
