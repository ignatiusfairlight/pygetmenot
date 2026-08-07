package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
	
	versionCmd = &cobra.Command{
	Use: "version",
	Short: "...",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

)

func init() {
	rootCmd.AddCommand(versionCmd)
}
