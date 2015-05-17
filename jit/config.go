package jit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/robhurring/jit/utils"
)

const (
	ConfigFilename = "config.json"
)

var (
	AppConfig *Config
)

func init() {
	AppConfig = GetConfig()
}

func ConfigPath() string {
	return path.Join(os.Getenv("HOME"), ".config", "jit")
}

type Config struct {
	Jira            *JiraConfig       `json:"jira"`
	Github          *GithubConfig     `json:"github"`
	MaxBranchLength int               `json:"maxBranchLength"`
	DefaultRemote   string            `json:"defaultRemote"`
	DefaultBranch   string            `json:"defaultBranch"`
	AssociatedPaths []string          `json:"associatedPaths"`
	UserMap         map[string]string `json:"userMap"`
}

type GithubConfig struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func (c *GithubConfig) FilledOut() bool {
	return (filledOut(c.Username) && filledOut(c.Token))
}

type JiraConfig struct {
	Host           string `json:"host"`
	ApiPath        string `json:"api_path"`
	ActivityPath   string `json:"activity_path"`
	Login          string `json:"login"`
	Password       string `json:"password"`
	DefaultProject string `json:"defaultProject"`
}

func (c *JiraConfig) FilledOut() bool {
	return (filledOut(c.Host) && filledOut(c.Login) && filledOut(c.Password))
}

func createConfig() {
	config := &Config{
		Jira: &JiraConfig{
			Host:         "",
			ApiPath:      "/rest/api/2",
			ActivityPath: "/activity",
		},
		Github: &GithubConfig{
			Username: "",
			Token:    "",
		},
		MaxBranchLength: 25,
		DefaultRemote:   "origin",
		DefaultBranch:   "dev",
		AssociatedPaths: []string{},
		UserMap:         make(map[string]string, 0),
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

func GetConfig() (config *Config) {
	configFile := path.Join(ConfigPath(), ConfigFilename)

	if !utils.FileExists(configFile) {
		createConfig()
	}

	config = new(Config)
	data, err := ioutil.ReadFile(configFile)
	utils.Check(err)

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

func SaveConfig() {
	WriteConfig(GetConfig())
}
