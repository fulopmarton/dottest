package cmd

import (
	mappingservice "dottest/internal/services"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
)

func getDefaultDomainName() string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Base(cwd)
}

var linkCmd = &cobra.Command{
	Use:   "link <port> [domain-name]",
	Short: "Run the dottest dameon server",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// var port uint
		domainName := getDefaultDomainName()
		if len(args) < 1 {
			log.Fatal("You must specify a port to link to the dottest daemon server!")
		}
		if len(args) >= 2 {
			domainName = args[1]
		}
		port, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("Invalid port number: %s", args[0])
		}
		fmt.Printf("Linking domain '%s' to port '%d'\n", domainName, port)
		err = mappingservice.AddMapping(domainName, fmt.Sprintf("http://localhost:%d", port))
		if err != nil {
			log.Fatalf("Failed to link domain '%s' to port '%d': %v\n", domainName, port, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(linkCmd)
}
