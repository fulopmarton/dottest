package main

import (
	"dottest/internal/dns"
	"dottest/internal/proxy"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	godotenv.Load(".env.local")
	go dns.StartDNSServer()
	proxy.StartReverseProxy()
}
