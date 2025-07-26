package cmd

import (
	"dottest/utils"
	"os"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Set up the project",
	Run: func(cmd *cobra.Command, args []string) {
		certificateFolder := utils.GetAppDataPath("./data/certs/")
		os.MkdirAll(certificateFolder, 0755)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
