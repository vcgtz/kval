/*
Copyright © 2023 Vicente Gutiérrez vicente@dotruntime.com
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vcgtz/local-store/internal/localstoreutil"
	"golang.design/x/clipboard"
	"os"
)

var copy bool

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [KEY] [FLAGS]",
	Short: "Get a single value stored",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value, _ := localstoreutil.GetValue(key)

		fmt.Println("\n", value)

		if copy {
			err := clipboard.Init()

			if err != nil {
				fmt.Println("An error occurs: ", err)
				os.Exit(1)
			}

			clipboard.Write(clipboard.FmtText, []byte(value))
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	getCmd.Flags().BoolVarP(&copy, "copy", "c", false, "copy the value to the clipboard.")
}
