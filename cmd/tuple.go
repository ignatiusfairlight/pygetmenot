package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	tupleCmd = &cobra.Command{
	Use: "tuple",
	Short: "...",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You can't change me!")
	},
}

)

func init() {
	rootCmd.AddCommand(tupleCmd)
}