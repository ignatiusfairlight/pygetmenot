package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	dictCmd = &cobra.Command{
	Use: "dict",
	Short: "...",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Python dictionary, yes.")
	},
}

)

func init() {
	rootCmd.AddCommand(dictCmd)
}