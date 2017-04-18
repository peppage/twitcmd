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

	"os"
	"text/tabwriter"

	"strings"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/fatih/color"
	wordwrap "github.com/mitchellh/go-wordwrap"
	"github.com/spf13/cobra"
)

// whoisCmd represents the whois command
var whoisCmd = &cobra.Command{
	Use:   "whois [screenname]",
	Short: "Profile information for a user",
	Long: `Full twitter profile information for a user in a table.Execute
	
Example:
twitcmd whois peppage`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			er("Incorrect amount of arguments")
		}

		user, _, err := client.Users.Show(&twitter.UserShowParams{
			ScreenName: args[0],
		})
		if err != nil {
			er("Getting user failed " + err.Error())
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.DiscardEmptyColumns)
		fmt.Fprintln(w, "ID\t"+user.IDStr)
		fmt.Fprintln(w, "Name\t"+user.Name)
		fmt.Fprintln(w, "Screen name\t"+user.ScreenName)
		if user.Protected {
			fmt.Fprintln(w, color.RedString("Protected"))
		}
		if user.Verified {
			fmt.Fprintln(w, color.BlueString("Verified"))
		}
		for i, s := range strings.Split(wordwrap.WrapString(user.Description, 60), "\n") {
			if i == 0 {
				fmt.Fprintln(w, "Bio\t"+s)
			} else {
				fmt.Fprintln(w, " \t"+s)
			}
		}

		fmt.Fprintln(w, "Location\t"+user.Location)
		fmt.Fprintln(w, "Email\t"+user.Email)
		fmt.Fprintln(w, "URL\t"+user.URL)
		fmt.Fprintf(w, "Likes\t%d\n", user.FavouritesCount)
		fmt.Fprintf(w, "Followers\t%d\n", user.FollowersCount)
		fmt.Fprintf(w, "Tweets\t%d\n", user.StatusesCount)
		fmt.Fprintln(w, "Created\t"+user.CreatedAt)
		w.Flush()
	},
}

func init() {
	RootCmd.AddCommand(whoisCmd)
}
