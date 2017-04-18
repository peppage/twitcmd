package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/fatih/color"
	wordwrap "github.com/mitchellh/go-wordwrap"
)

func er(msg interface{}) {
	fmt.Fprintln(os.Stderr, "Error:", msg)
	os.Exit(-1)
}

func newFalse() *bool {
	b := false
	return &b
}

func newTrue() *bool {
	b := true
	return &b
}

func parseID(id string) int64 {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		er("That tweet id is not a number")
	}
	return i
}

func displayTweets(tweets []twitter.Tweet, reverse bool) {
	for i := range tweets {
		t := tweets[i]
		if reverse {
			t = tweets[len(tweets)-1-i]
		}

		color.Yellow("@" + t.User.ScreenName)
		color.Green(strconv.FormatInt(t.ID, 10))
		color.White(wordwrap.WrapString(t.Text, 80))
		fmt.Println("")
	}
}
