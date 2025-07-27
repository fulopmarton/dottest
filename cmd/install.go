package cmd

import (
	"dottest/utils"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Set up the project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Installing dottest...")
		certificateFolder := utils.GetAppDataPath("./data/certs/")
		os.MkdirAll(certificateFolder, 0777)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
