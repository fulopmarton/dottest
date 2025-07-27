package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Command for testing purposes",
	Run: func(cmd *cobra.Command, args []string) {
		res, _ := exec.Command("whoami").CombinedOutput()
		fmt.Printf("%s", res)
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
