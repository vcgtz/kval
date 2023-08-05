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

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Delete all the exising keys",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := localstoreutil.Clean()

		if err != nil {
			fmt.Println("An error occurs: ", err)
			os.Exit(1)
		}

		msgFormat := color.New(color.FgRed, color.Bold)
		_, _ = msgFormat.Println("\nAll the keys were removed successfully.")
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cleanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cleanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
