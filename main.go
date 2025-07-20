package main

import (
	"dottest/internal/dns"
	"dottest/internal/proxy"
)

func main() {
	go dns.StartDNSServer()
	proxy.StartReverseProxy()
}
