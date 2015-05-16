                                   __     __     ______
                                  /\ \   /\ \   /\__  _\
                                 _\_\ \  \ \ \  \/_/\ \/
                                /\_____\  \ \_\    \ \_\
                                \/_____/   \/_/     \/_/
                                              Jira + Git

###### About

Jerry is a small command-line tool used to aid my day-to-day workflow using jira
and github. I'm not sure how useful it will be for anyone else not working the same way.

This is a go re-write of my node tool [jerry](https://github.com/robhurring/jerry)

## Installation

Make sure you have a working Go environment (go 1.1 is required). See the install instructions.

To install jit, simply run:

`$ go get github.com/robhurring/jit`

Make sure your PATH includes to the $GOPATH/bin directory so your commands can be easily used:

`export PATH=$PATH:$GOPATH/bin`

## Commands

#### branch, br
###### In Progress...

The branch command will take a TICKET_ID param and create a named branch for you in your current repo.

#### pull-request, pr
###### In Progress...

If you are on a feature branch Jerry can help you open a fleshed-out pull-request. If you are the lazy-type, this is really the best feature.

Jit will detect the ticket from your branch name, lookup some basic details about the ticket from Jira and use that info to build out a pull-request. If all goes well, it will open a browser window to your new PR so you can do some adjustments.

#### info, i
###### In Progress...

Print some basic information about a given ticket. If not ticket is passed in through the args list jerry will check your current branch and extract the ticket from there.

#### open, o
###### In Progress...

Open the given ticket in the browser. Jit will extract the ticket ID from the feature branch if possible and no ID was given in the args.

#### copy, cp
###### In Progress...

Copy the given ticket to the clipboard. Jit will extract the ticket ID from the feature branch if possible and no ID was given in the args.


### Credits

I lovingly stole a bunch of functions from [Github's hub](https://github.com/github/hub) repo.
