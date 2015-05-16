package jit

import (
	"fmt"

	"github.com/google/go-github/github"
)

func FindUsername(fullName string) (username string) {
	username = fullName
	client := NewGithubClient()

	options := &github.SearchOptions{}
	criteria := fmt.Sprintf("fullname:%s", fullName)

	results, _, err := client.Search.Users(criteria, options)
	if err != nil {
		return
	}

	// there can only be one
	if len(results.Users) == 1 {
		username = fmt.Sprintf("@%s", *results.Users[0].Login)
	}

	return
}
