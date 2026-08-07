package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	_ "embed"
)

//go:embed notes/loop.md
var loopNotes string

var (
	loopCmd = &cobra.Command{
	Use: "loop",
	Short: "...",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(loopNotes)
	},
}

)

func init() {
	rootCmd.AddCommand(loopCmd)
}