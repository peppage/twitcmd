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

	"github.com/dghubble/go-twitter/twitter"
	"github.com/mitchellh/go-wordwrap"
	"github.com/spf13/cobra"
)

// deletestatusCmd represents the deletestatus command
var deletestatusCmd = &cobra.Command{
	Use:   "status [tweet id...]",
	Short: "Delete a tweet",
	Long: `Delete your tweet from twitter

Example:
twitcmd delete status 784103823405178881
	`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			er("Missing arguments")
		}

		for i := range args {
			id := parseID(args[i])

			t, _, err := client.Statuses.Destroy(id, &twitter.StatusDestroyParams{
				TrimUser: newTrue(),
			})
			if err == nil {
				fmt.Println("@ deleted the Tweet " + wordwrap.WrapString(t.Text, 80))
			}
		}
	},
}

func init() {
	deleteCmd.AddCommand(deletestatusCmd)
}
