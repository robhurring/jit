package jit

import (
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var githubConfig *GithubConfig

func init() {
	githubConfig = AppConfig.Github
}

func NewAuthenticatedGithubClient() *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubConfig.Token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	return github.NewClient(tc)
}

func NewGithubClient() *github.Client {
	return github.NewClient(nil)
}
