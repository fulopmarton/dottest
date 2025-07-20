package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dottest",
	Short: "A tool for testing local websites without custom domain names",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to dottest! Use 'dottest daemon' to run the server.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}
