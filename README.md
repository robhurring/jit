                                   __     __     ______
                                  /\ \   /\ \   /\__  _\
                                 _\_\ \  \ \ \  \/_/\ \/
                                /\_____\  \ \_\    \ \_\
                                \/_____/   \/_/     \/_/
                                              Jira + Git


## About

JIT is a small command-line tool used to aid my day-to-day workflow using jira
and github. I'm not sure how useful it will be for anyone else not working the same way.

This is a go re-write of my node tool [jerry](https://github.com/robhurring/jerry)

## Installation

Make sure you have a working Go environment (go 1.1 is required). See the install instructions.

To install jit, download the latest release:

* [Latest Release](https://github.com/robhurring/jit/releases/latest)

To install jit from source:

`$ go get github.com/robhurring/jit`

Make sure your PATH includes to the $GOPATH/bin directory so your commands can be easily used:

`export PATH=$PATH:$GOPATH/bin`

## Setup

To configure JIT so it can talk with Github and Jira, you need to add the config file to `~/.config/jit/config.json`

An example configuration looks like this:

```js
// ~/.config/jit/config.json
{
  "jira": {
    "host": "https://mycompany.atlassian.net",
    "api_path": "/rest/api/2",
    "activity_path": "/activity",
    "login": "JIRA_PASSWORD",
    "password": "JIRA_PASSWORD",
    // when using naked issue numbers with JIT, this project will be pre-pended
    "defaultProject": ""
  },
  "github": {
    "username": "GITHUB_USERNAME",
    "token": "TOKEN"
  },
  // when generating branch names from issue summaries, this is the max length
  // it will be. truncated at the last full word
  "maxBranchLength": 35,
  // when generating a pull-request any repo in this list of paths with a branch
  // matching the current repo's branch will be listed as "associated"
  "associatedPaths": [
    "~/Code"
  ],
  // when generating a pull-request JIT will attempt to map names from JIRA to
  // github usernames using the github API. any name listed here will override
  // that.
  "userMap": {
    "rob hurring": "robhurring"
  }
}
```

## Commands

#### branch, br

Create a new branch for the given ISSUE. If the branch already exists JIT will issue a `checkout` instead.


  * `--preview, -p`  Preview branch name
  * `--copy, -c`   Copy branch name  (OSX only)

###### Example

```sh
$ jit branch DEV-12
# => git checkout -b DEV-12_update_the_wizzy_wigs
```

#### pull-request, pr

Create a new pull-request for the given ISSUE. JIT will try to map the issue's `CodeReviewer`'s full name against the Github API - if only 1 result is found it will be used. If the name cannot be found (or multiple results were returned), JIT will just use the full-name.

__HINT:__ For common names or people who don't use their full-name in github, add them to the `userMap` settings hash.

  * `--preview, -p`  Preview the pull-request
  * `--copy, -c`   Copy the pull-request body to the clipboard (OSX only)

###### Example

```sh
$ jit pull-reqest DEV-12
# => creates a pull-request for the current branch referencing DEV-12

$ jit pull-reqest --preview DEV-12

# DEV-12: test
#
# /cc @robhurring
#
# [JIRA DEV-12](https://mycompany.atlassian.net/rest/api/2/issue/90210): Test
#
# ### Associated
#
# some-other-repo
#
# ### Summary
#
# * Changed A, B, C
#
# ### Testing
#
# `rake spec`
```

#### info, in

Print some basic information about a given issue.

###### Example

```sh
$ jit info DEV-12

# DEV-12: Test
# https://mycompany.atlassian.net/rest/api/2/issue/90210
#
# Creator:  Rob Hurring
# Developer: Rob Hurring
# Reviewer: Rob Hurring
# Assigned: Rob Hurring
#
# -----------------------8<-------------------------------------------------------
#
# Links (2):
#
#   Blocks
#   DEV-15: [New]:  Do something relating to this ticket.
#
#   Relates
#   IT-52: [In Production]:  Some IT issue related to this ticket.
#
# -----------------------8<-------------------------------------------------------
#
# Status: New
#
# Testing. This is my issue description
#
# -----------------------8<-------------------------------------------------------
#
# Comments (2):
#
# "test"
# Reginald Sombernotch
#
# "hello world!"
# Drake Thunderfist
```

#### open, o

Open the given ticket in the browser. (OSX only. Using the `open` command currently.)

```sh
$ jit open DEV-12
# => opens the issue in the browser
```

#### copy, cp

Copy the given issue URL to the clipboard. (OSX only.)

```sh
$ jit copy DEV-12
# => Copied! https://mycompany.atlassian.net/browse/DEV-12
```

## Contributing

1. Fork it ( https://github.com/robhurring/jit/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request

## Credits

I :heart: stole a bunch of functions from [Github's hub](https://github.com/github/hub) repo.
