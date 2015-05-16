package git

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/robhurring/jit/cmd"
)

var (
	protocolRe = regexp.MustCompile("^[a-zA-Z_-]+://")
)

type Remote struct {
	Name string
	URL  *url.URL
}

type GithubProject struct {
	Owner string
	Name  string
}

func (r *Remote) Project() (project *GithubProject, err error) {
	parts := strings.SplitN(r.URL.Path, "/", 4)
	if len(parts) <= 2 {
		err = fmt.Errorf("Invalid GitHub URL: %s", r.URL)
		return
	}

	name := strings.TrimSuffix(parts[2], ".git")
	project = &GithubProject{
		Owner: parts[1],
		Name:  name,
	}

	return
}

func ParseRemote(rawURL string) (u *url.URL, err error) {
	if !protocolRe.MatchString(rawURL) &&
		strings.Contains(rawURL, ":") &&
		// not a Windows path
		!strings.Contains(rawURL, "\\") {
		rawURL = "ssh://" + strings.Replace(rawURL, ":", "/", 1)
	}

	u, err = url.Parse(rawURL)
	if err != nil {
		return
	}

	if u.Scheme != "ssh" {
		return
	}

	return
}

func OriginRemote() (remote Remote, err error) {
	remotes, err := Remotes()
	if err != nil {
		return
	}

	for _, r := range remotes {
		if r.Name == "origin" {
			remote = r
			break
		}
	}

	return
}

func Remotes() (remotes []Remote, err error) {
	re := regexp.MustCompile(`(.+)\s+(.+)\s+\((push|fetch)\)`)

	output, err := cmd.New("git").WithArgs("remote", "-v").CombinedOutput()
	if err != nil {
		return
	}

	// build the remotes map
	rs := strings.Split(string(output), "\n")
	remotesMap := make(map[string]string)
	for _, r := range rs {
		if re.MatchString(r) {
			match := re.FindStringSubmatch(r)
			name := strings.TrimSpace(match[1])
			url := strings.TrimSpace(match[2])
			remotesMap[name] = url
		}
	}

	// the rest of the remotes
	for n, u := range remotesMap {
		url, e := ParseRemote(u)
		if e == nil {
			remotes = append(remotes, Remote{Name: n, URL: url})
		}
	}

	return
}
