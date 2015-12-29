package git

import (
	"strings"

	"github.com/robhurring/cmd"
)

type fileMatcher func(string) bool

func ModifiedFiles(base string) ([]string, error) {
	output, err := cmd.New("git").WithArgs("diff", "--name-only", base).CombinedOutput()
	if err != nil {
		return nil, err
	}

	list := strings.Split(string(output), "\n")

	return list, nil
}

func ModifiedMatching(base string, matcher fileMatcher) ([]string, error) {
	matches := make([]string, 0)

	list, err := ModifiedFiles(base)
	if err != nil {
		return nil, err
	}

	for _, filename := range list {
		if match := matcher(filename); match {
			matches = append(matches, filename)
		}
	}

	return matches, nil
}

func ModifiedRailsFiles(base string) ([]string, error) {
	return ModifiedMatching(base, railsFileMatcher)
}

func ModifiedSpecFiles(base string) ([]string, error) {
	return ModifiedMatching(base, specFileMatcher)
}

func specFileMatcher(filename string) (matched bool) {
	paths := []string{
		"_spec.rb",
	}

	for _, path := range paths {
		if strings.Contains(filename, path) && strings.HasSuffix(filename, ".rb") {
			matched = true
			return
		}
	}

	return false
}

func railsFileMatcher(filename string) (matched bool) {
	paths := []string{
		"app/",
	}

	for _, path := range paths {
		if strings.Contains(filename, path) && strings.HasSuffix(filename, ".rb") {
			matched = true
			return
		}
	}

	return false
}
