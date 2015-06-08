package commands

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/robhurring/jit/cmd"
	"github.com/robhurring/jit/git"
	"github.com/robhurring/jit/ui"
)

func init() {
	CmdRunner.Add(&cli.Command{
		Name:      "hub",
		ShortName: "h",
		Usage:     "Information about GitHub",
		Subcommands: []cli.Command{
			{
				Name:  "branch",
				Usage: "Default GitHub branch",
				Action: func(c *cli.Context) {
					branch, err := git.DefaultBranch()
					if err != nil {
						panic(err)
					}
					fmt.Println(branch.ShortName())
				},
			},
			{
				Name:  "owner",
				Usage: "GitHub repo owner",
				Action: func(c *cli.Context) {
					project, err := getGitProject()
					if err != nil {
						panic(err)
					}
					fmt.Println(project.Owner)
				},
			},
			{
				Name:  "name",
				Usage: "GitHub repo name",
				Action: func(c *cli.Context) {
					project, err := getGitProject()
					if err != nil {
						panic(err)
					}
					fmt.Println(project.Name)
				},
			},
			{
				Name:  "url",
				Usage: "GitHub repo URL",
				Action: func(c *cli.Context) {
					url, err := githubURL()
					if err != nil {
						panic(err)
					}

					fmt.Println(url)
				},
			},
			{
				Name:  "open",
				Usage: "Open repo in GitHub",
				Action: func(c *cli.Context) {
					url, err := githubURL()
					if err != nil {
						panic(err)
					}

					ui.Printf("@{!w}Opening@| %s\n", url)
					cmd.Open(url)
				},
			},
		},
	})
}

func getGitProject() (project *git.GithubProject, err error) {
	origin, err := git.OriginRemote()
	if err != nil {
		return
	}

	project, err = origin.Project()
	if err != nil {
		return
	}

	return
}

func githubURL() (string, error) {
	project, err := getGitProject()
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://github.com/%s/%s.git", project.Owner, project.Name)

	return url, nil
}
