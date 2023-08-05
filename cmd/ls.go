/*
Copyright © 2023 Vicente Gutiérrez vicente@dotruntime.com
*/
package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/vcgtz/local-store/internal/localstoreutil"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all the stores keys",
	Run: func(cmd *cobra.Command, args []string) {
		keys, _ := localstoreutil.GetKeys()

		if len(keys) == 0 {
			fmt.Println("\nNo stored keys.")
			return
		}

		fmt.Print("\n")
		titleFormat := color.New(color.Bold)
		_, _ = titleFormat.Println("Keys")

		for _, key := range keys {
			fmt.Printf("-> ")
			keyFormat := color.New(color.FgYellow)
			_, _ = keyFormat.Println(key)
		}

		fmt.Print("\n")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
