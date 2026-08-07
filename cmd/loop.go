package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	loopCmd = &cobra.Command{
	Use: "loop",
	Short: "...",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Praise Oroboros!")
	},
}

)

func init() {
	rootCmd.AddCommand(loopCmd)
}