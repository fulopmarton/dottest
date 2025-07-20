package cmd

import (
	"dottest/config"
	"dottest/internal/dns"
	"dottest/internal/proxy"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Run the dottest dameon server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting dottest daemon server...")

		err := godotenv.Load()
		log.Printf("Default config: %v", config.DefaultConfig)
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		godotenv.Load(".env.local")
		go dns.StartDNSServer()
		proxy.StartReverseProxy()
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}
