package commands

import (
	"strings"

	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
	"github.com/robhurring/jit/cmd"
	"github.com/robhurring/jit/git"
	"github.com/robhurring/jit/jit"
	"github.com/robhurring/jit/ui"
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
			cli.BoolFlag{
				Name:  "preview, p",
				Usage: "Preview branch name",
			},
			cli.BoolFlag{
				Name:  "copy, c",
				Usage: "Copy branch name",
			},
		},
		Action: func(c *cli.Context) {
			// panic unless we have a .git dir
			if _, err := git.Dir(); err != nil {
				panic(err)
			}

			currentBranch, err := git.CurrentBranch()
			if err != nil {
				panic(err)
			}
			head := currentBranch.Name

			base := c.String("base")
			if base == "" {
				branch, err := git.DefaultBranch()
				if err != nil {
					panic(err)
				}
				base = branch.ShortName()
			}

			key, err := DetectIssue(c.Args())
			if err != nil {
				panic(err)
			}

			issue, err := jit.GetIssue(key, true)
			if err != nil {
				panic(err)
			}

			pull := makePull(head, base, issue)

			if c.Bool("preview") {
				ui.PrintTemplate("pull-request.info", pull)
				return
			}

			if c.Bool("copy") {
				output := ui.RenderTemplate("pull-request.info", pull)
				if err := cmd.Copy(output); err != nil {
					panic(err)
				} else {
					ui.Printf("@{!w}Copied!@|\n")
				}
				return
			}

			createPullRequest(pull)
		},
	})
}

type pullRequestTemplateData struct {
	CodeReviewer string
	Key          string
	URL          string
	Title        string
	Associated   []string
}

func pullRequestTemplate(issue *jit.Issue) *pullRequestTemplateData {
	associated := git.AssociatedProjects(issue.BranchName())
	data := &pullRequestTemplateData{
		CodeReviewer: issue.Fields.CodeReviewer.DisplayName,
		Key:          issue.Key,
		URL:          issue.Self,
		Title:        issue.Fields.Summary,
		Associated:   associated,
	}

	return data
}

func makePull(head, base string, issue *jit.Issue) *github.NewPullRequest {
	body := ui.RenderTemplate("pull-request.body", pullRequestTemplate(issue))
	body = strings.TrimSpace(body)
	title := issue.String()

	pull := &github.NewPullRequest{
		Head:  &head,
		Base:  &base,
		Body:  &body,
		Title: &title,
	}

	return pull
}

func createPullRequest(pull *github.NewPullRequest) {
	origin, err := git.OriginRemote()
	if err != nil {
		panic(err)
	}

	project, err := origin.Project()
	if err != nil {
		panic(err)
	}

	client := jit.NewAuthenticatedGithubClient()
	newPull, _, err := client.PullRequests.Create(project.Owner, project.Name, pull)
	if err != nil {
		panic(err)
	}

	ui.Printf("@{!w}Created!@| %s\n", *newPull.HTMLURL)
	cmd.Open(*newPull.HTMLURL)
}
