/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vcgtz/local-store/internal/localstoreutil"
)

var force bool

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set [KEY] [VALUE]",
	Short: "Store a single text value.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value := args[1]

		successMsg, err := localstoreutil.StoreValue(key, value, force)

		if err != nil {
			fmt.Println("An error occurs: ", err)
			os.Exit(1)
		}

		fmt.Println("\n", successMsg)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	setCmd.Flags().BoolVarP(&force, "force", "f", false, "force to overwrite an existing value.")
}
