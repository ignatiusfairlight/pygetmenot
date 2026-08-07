package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
	Use: "list",
	Short: "...",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Python's list. We are definitely not an array, ahahaha...")
	},
}

)

func init() {
	rootCmd.AddCommand(listCmd)
}