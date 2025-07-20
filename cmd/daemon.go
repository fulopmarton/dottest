package cmd

import (
	"dottest/internal/dns"
	"dottest/internal/proxy"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Run the dottest dameon server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting dottest daemon server...")

		godotenv.Load(".env.local")
		go dns.StartDNSServer()
		proxy.StartReverseProxy()
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}
