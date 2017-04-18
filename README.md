# twitcmd
[![Go Report Card](https://goreportcard.com/badge/github.com/peppage/twitcmd)](https://goreportcard.com/report/github.com/peppage/twitcmd)

#### twitcmd is a command-line tool for accessing Twitter.
Very simple way to read twitter and respond from the command line.

## Install
    > go install github.com/peppage/twitcmd

## Configuration
Twitter API requires access keys. Right now to get out of doing oauth I create a new Twitter application and create
a user key for that application. It's a pretty easy process, go here https://apps.twitter.com/app/new. You can see the
required keys in the .twitcmd.yaml.example file that should be in your $HOME dir unless you want to specifically send
the location.

## Help
Help is available for all commands. If you need help for a subcommand type the command after help.

    twitcmd help

#### Update status
    twitcmd update "Your message that you want to tweet"

#### Timeline with 10 tweets reversed
    twitcmd timeline -n 10 -r

#### Reply to someone
    twitcmd reply 852174540075671556 "Your reply message"

#### Delete tweets
    twitcmd delete 852174540075671556 853926060496375809

#### Get mentions
    twitcmd mentions -n 10

#### Lookup a user
    twitcmd whois Peppage