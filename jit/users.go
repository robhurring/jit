package jit

import (
	"fmt"
	"strings"

	"github.com/google/go-github/github"
)

func FindUsername(fullName string) (username string) {
	username = fullName

	if ok, configUsername := lookupConfigUsername(fullName); ok {
		username = configUsername
		return
	}

	if ok, githubUsername := loopupGithubUsername(fullName); ok {
		username = githubUsername
		return
	}

	return
}

func lookupConfigUsername(name string) (ok bool, username string) {
	ok = false

	config := GetConfig()
	userMap := config.UserMap

	if userMap == nil {
		ok = false
		return
	}

	search := strings.ToLower(name)
	for n, un := range userMap {
		if n == search {
			ok = true
			username = fmt.Sprintf("@%s", un)
			break
		}
	}

	return
}

func loopupGithubUsername(name string) (ok bool, username string) {
	ok = false
	client := NewGithubClient()

	options := &github.SearchOptions{}
	criteria := fmt.Sprintf("fullname:%s", name)

	results, _, err := client.Search.Users(criteria, options)
	if err != nil {
		panic(err)
		return
	}

	// there can only be one
	if len(results.Users) == 1 {
		username := fmt.Sprintf("@%s", *results.Users[0].Login)
		return true, username
	}

	return
}
