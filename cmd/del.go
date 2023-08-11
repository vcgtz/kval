/*
Copyright © 2023 Vicente Gutiérrez vicente@dotruntime.com
*/
package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/vcgtz/local-store/internal/localstoreutil"
	"os"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "del [KEY]",
	Short: "Delete an existing key",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]

		status, err := localstoreutil.DeleteValue(key)

		if err != nil {
			fmt.Println("An error occurs: ", err)
			os.Exit(1)
		}

		if status == "no_exist" {
			fmt.Print("The key")

			keyFormat := color.New(color.FgYellow, color.Bold)
			_, _ = keyFormat.Printf(" %s ", key)

			fmt.Println("does not exist")
		}

		if status == "success" {
			fmt.Print("The key")

			keyFormat := color.New(color.FgYellow, color.Bold)
			_, _ = keyFormat.Printf(" %s ", key)

			fmt.Println("was deleted")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
