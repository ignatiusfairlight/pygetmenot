package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
	Use: "pygetmenot",
	Short: "A terminal app that reminds you how to use Python features.",
	Long: `A terminal app designed to assist in reminding you how to write Python
	codes with detailed explanation and code examples enough for you to understand 
	and continue your work when you are stuck in writing any Python projects.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is Pygetmenot. A terminal app to assist you in writing Python codes.")
	},
}

)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}