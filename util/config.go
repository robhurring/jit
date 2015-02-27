package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

const (
	ConfigFilename = "config.json"
)

func ConfigPath() string {
	return path.Join(os.Getenv("HOME"), ".config", "jit")
}

type Config struct {
	Jira            *JiraConfig   `json:"jira"`
	Github          *GithubConfig `json:"github"`
	MaxBranchLength int           `json:"maxBranchLength"`
	DefaultRemote   string        `json:"defaultRemote"`
	DefaultBranch   string        `json:"defaultBranch"`
	AssociatedPaths []string      `json:"associatedPaths"`
}

type GithubConfig struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func (c *GithubConfig) FilledOut() bool {
	return (filledOut(c.Username) && filledOut(c.Token))
}

type JiraConfig struct {
	Host         string `json:"host"`
	ApiPath      string `json:"api_path"`
	ActivityPath string `json:"activity_path"`
	Login        string `json:"login"`
	Password     string `json:"password"`
}

func (c *JiraConfig) FilledOut() bool {
	return (filledOut(c.Host) && filledOut(c.Login) && filledOut(c.Password))
}

func createConfig() {
	jiraConfig := &JiraConfig{
		Host:         "",
		ApiPath:      "/rest/api/2",
		ActivityPath: "/activity",
	}

	config := &Config{
		Jira:            jiraConfig,
		MaxBranchLength: 25,
		DefaultRemote:   "origin",
		DefaultBranch:   "dev",
		AssociatedPaths: []string{},
	}
	WriteConfig(config)
}

func filledOut(value string) bool {
	return len(value) > 0
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func GetConfig() (config *Config, err error) {
	configFile := path.Join(ConfigPath(), ConfigFilename)

	if !FileExists(configFile) {
		createConfig()
	}

	config = new(Config)
	data, err := ioutil.ReadFile(configFile)
	json.Unmarshal(data, config)

	return
}

func WriteConfig(c *Config) {
	configPath := ConfigPath()
	configFile := path.Join(configPath, ConfigFilename)

	err := os.MkdirAll(configPath, 0755)
	check(err)

	data, jerr := json.MarshalIndent(c, "", "  ")
	check(jerr)

	ioerr := ioutil.WriteFile(configFile, data, 0644)
	check(ioerr)
}