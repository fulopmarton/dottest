package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Run the dottest dameon server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting dottest daemon server...")
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)
}
