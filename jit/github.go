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

// func CreatePullRequest(pull *github.NewPullRequest) (*github.PullRequest, error) {
// Create(owner string, repo string, pull *NewPullRequest) (*PullRequest, *Response, error) {

// }

// type NewPullRequest struct {
//   Title *string `json:"title,omitempty"`
//   Head  *string `json:"head,omitempty"`
//   Base  *string `json:"base,omitempty"`
//   Body  *string `json:"body,omitempty"`
//   Issue *int    `json:"issue,omitempty"`
// }
