// Copyright © 2017 Michael Peppler <peppage@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"strconv"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/fatih/color"
	"github.com/mitchellh/go-wordwrap"
	"github.com/spf13/cobra"
)

var timelineCount int32
var timelineReverse bool

// timelineCmd represents the timeline command
var timelineCmd = &cobra.Command{
	Use:   "timeline [flags]",
	Short: "Your home timeline",
	Long: `The twitter home timeline.

Show tweets users are tweeting right now!

Example:
To show the last 5 tweets

twitcmd timeline -n 5`,
	Run: func(cmd *cobra.Command, args []string) {
		tweets, _, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
			Count: int(timelineCount),
		})
		if err != nil {
			er("Getting timeline failed " + err.Error())
		}

		for i := range tweets {
			t := tweets[i]
			if timelineReverse {
				t = tweets[len(tweets)-1-i]
			}

			color.Yellow("@" + t.User.ScreenName)
			color.Green(strconv.FormatInt(t.ID, 10))
			color.White(wordwrap.WrapString(t.Text, 80))
			fmt.Println("")
		}
	},
}

func init() {
	RootCmd.AddCommand(timelineCmd)

	timelineCmd.Flags().Int32VarP(&timelineCount, "number", "n", 200, "The number of tweets to get")
	timelineCmd.Flags().BoolVarP(&timelineReverse, "reverse", "r", false, "Reverse the sort order")
}
