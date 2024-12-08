/*
Copyright © 2024 Sarvsav Sharma <sarvsav+github@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"

	"github.com/sarvsav/go-mongodb/internals"
	"github.com/sarvsav/go-mongodb/models"
	"github.com/spf13/cobra"
)

func WithLongListing(b bool) internals.OptionsLsFunc {
	return func(c *models.LsOptions) error { c.LongListing = b; return nil }
}

func WithColor(color bool) internals.OptionsLsFunc {
	return func(c *models.LsOptions) error { c.Color = color; return nil }
}

func WithArgs(args []string) internals.OptionsLsFunc {
	return func(c *models.LsOptions) error { c.Args = args; return nil }
}

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		longListingValue, _ := cmd.Flags().GetBool("long")
		colorValue, _ := cmd.Flags().GetBool("color")
		// Use the longListingValue to determine if long listing is required
		fmt.Println("ls called with long listing value:", longListingValue)
		internals.Ls(WithLongListing(longListingValue), WithColor(colorValue), WithArgs(args))
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	lsCmd.Flags().BoolP("long", "l", false, "Long listing format of databases and collections")
	lsCmd.Flags().BoolP("color", "c", false, "Add colors to the output")
}
