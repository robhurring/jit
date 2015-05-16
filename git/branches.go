package git

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/robhurring/jit/cmd"
)

type Branch struct {
	Name string
}

func (b *Branch) ShortName() string {
	reg := regexp.MustCompile("^refs/(remotes/)?.+?/")
	return reg.ReplaceAllString(b.Name, "")
}

func (b *Branch) Exists() (exists bool, err error) {
	exists = false

	branches, err := BranchList()
	if err != nil {
		return
	}

	for _, branch := range branches {
		if branch.Name == b.Name {
			exists = true
			break
		}
	}

	return
}

func (b *Branch) Checkout() (string, error) {
	return cmd.New("git").WithArgs("checkout", b.Name).CombinedOutput()
}

func (b *Branch) Create() (string, error) {
	return cmd.New("git").WithArgs("branch", b.Name).CombinedOutput()
}

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

// func BranchExists(name string) (exists bool, err error) {
// 	exists = false

// 	branches, err := BranchList()
// 	if err != nil {
// 		return
// 	}

// 	for _, branch := range branches {
// 		if branch == name {
// 			exists = true
// 			break
// 		}
// 	}

// 	return
// }

// func Checkout(branch string) (string, error) {
// 	return cmd.New("git").WithArgs("checkout", branch).CombinedOutput()
// }

// func CreateBranch(name string) (string, error) {
// 	return cmd.New("git").WithArgs("branch", name).CombinedOutput()
// }

func CurrentBranch() (branch *Branch, err error) {
	output, err := cmd.New("git").WithArgs("rev-parse", "--abbrev-ref", "HEAD").CombinedOutput()
	name := strings.TrimSpace(output)
	branch = &Branch{Name: name}

	return
}

func BranchAtRef(paths ...string) (branch *Branch, err error) {
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
		name := strings.TrimPrefix(n, refPrefix)
		name = strings.TrimSpace(name)
		branch = &Branch{Name: name}
	} else {
		err = fmt.Errorf("No branch info in %s: %s", path, n)
	}

	return
}

func DefaultBranch() (branch *Branch, err error) {
	branch, err = BranchAtRef("refs", "remotes", "origin", "HEAD")
	if err != nil {
		return
	}

	if branch.Name == "" {
		branch = &Branch{
			Name: "refs/heads/master",
		}
	}

	return
}

// func ShortName(branch string) string {
// 	reg := regexp.MustCompile("^refs/(remotes/)?.+?/")
// 	return reg.ReplaceAllString(branch, "")
// }

func BranchList() (branches []*Branch, err error) {
	branchSet := make(map[string]int)

	branchCmd := cmd.New("git").WithArgs("branch", "-a", "--no-color")
	cutCmd := cmd.New("cut").WithArgs("-c", "3-")

	output, _, err := cmd.Pipeline(branchCmd, cutCmd)
	if err != nil {
		fmt.Println("erro.")
		return
	}

	list := strings.Split(string(output), "\n")
	for _, branch := range list {
		if normalized, ok := normalizeBranchName(branch); ok {
			if _, exists := branchSet[normalized]; !exists {
				branchSet[normalized] = 0
			}
		}
	}

	for name := range branchSet {
		b := &Branch{Name: name}
		branches = append(branches, b)
	}

	return branches, err
}
