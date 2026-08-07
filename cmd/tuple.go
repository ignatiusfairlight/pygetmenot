package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	_ "embed"
)

//go:embed notes/tuple.md
var tupleNotes string

var (
	tupleCmd = &cobra.Command{
	Use: "tuple",
	Short: "...",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(tupleNotes)
	},
}

)

func init() {
	rootCmd.AddCommand(tupleCmd)
}