package commands

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
	"github.com/robhurring/go-jira-client"
	"github.com/robhurring/jit/git"
)

func init() {
	CmdRunner.Add(&cli.Command{
		Name:      "pull-request",
		ShortName: "pr",
		Usage:     "Create a pull-request for the issue",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "base, b",
				Usage: "Pull-request base branch",
			},
		},

		Action: func(c *cli.Context) {
			// panic unless we have a .git dir
			// if _, err := git.Dir(); err != nil {
			// 	panic(err)
			// }

			// head, err := git.CurrentBranch()
			// if err != nil {
			// 	panic(err)
			// }

			// base := c.String("base")
			// if base == "" {
			// 	name, err := git.DefaultBranch()
			// 	if err != nil {
			// 		panic(err)
			// 	}
			// 	base = git.ShortName(name)
			// }

			// u, err := git.ParseRemote("git@github.com:lendkey/otto")
			// if err != nil {
			// 	panic(err)
			// }

			// project, err := git.GithubProjectFromURL(u)
			// if err != nil {
			// 	panic(err)
			// }

			// fmt.Println(project)
			origin, _ := git.OriginRemote()
			project, _ := origin.Project()

			fmt.Println(project.Name)
			// if err != nil {
			// 	panic(err)
			// }

			// fmt.Println(remote)

			// key, err := jit.FindIssueKey(c.Args())
			// if err != nil {
			// 	panic(err)
			// }

			// issue := jit.GetIssue(key, true)
			// pull := makePull(head, base, issue)
			// utils.Debug(pull)
		},
	})
}

func makePull(head, base string, issue gojira.Issue) *github.NewPullRequest {
	pull := &github.NewPullRequest{
		Head: &head,
		Base: &base,
	}

	return pull
}

// func pullBody(issue *jit.Issue) string {
// type NewPullRequest struct {
//   Title *string `json:"title,omitempty"`
//   Head  *string `json:"head,omitempty"`
//   Base  *string `json:"base,omitempty"`
//   Body  *string `json:"body,omitempty"`
//   Issue *int    `json:"issue,omitempty"`
// }

// }
