package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	_ "embed"
)

//go:embed notes/list.md
var listNotes string

var (
	listCmd = &cobra.Command{
	Use: "list",
	Short: "...",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(listNotes)
	},
}

)

func init() {
	rootCmd.AddCommand(listCmd)
}