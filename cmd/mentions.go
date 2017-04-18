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
	"github.com/dghubble/go-twitter/twitter"
	"github.com/spf13/cobra"
)

var mentionsCount int32
var mentionsReverse bool

// mentionsCmd represents the mentions command
var mentionsCmd = &cobra.Command{
	Use:   "mentions",
	Short: "The most recent tweets mentioning you",
	Long: `Your twitter mentions timeline.
	
Shows users interacting with you right now!

Example:
To show the last 5 mentions

twitcmd mentions -n 5`,
	Run: func(cmd *cobra.Command, args []string) {
		mentions, _, err := client.Timelines.MentionTimeline(&twitter.MentionTimelineParams{
			Count: int(mentionsCount),
		})
		if err != nil {
			er("Getting mention timeline failed " + err.Error())
		}

		displayTweets(mentions, mentionsReverse)
	},
}

func init() {
	RootCmd.AddCommand(mentionsCmd)

	mentionsCmd.Flags().Int32VarP(&mentionsCount, "number", "n", 200, "The number of mentions to get")
	mentionsCmd.Flags().BoolVarP(&mentionsReverse, "reverse", "r", false, "Reverse the sort order")
}
