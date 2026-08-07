package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	_ "embed"
)

//go:embed notes/dict.md
var dictNotes string

var (
	dictCmd = &cobra.Command{
	Use: "dict",
	Short: "...",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(dictNotes)
	},
}

)

func init() {
	rootCmd.AddCommand(dictCmd)
}