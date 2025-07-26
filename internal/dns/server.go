package dns

import (
	"dottest/config"
	mappingservice "dottest/internal/services"
	"fmt"
	"log"
	"strings"

	"github.com/miekg/dns"
)

func StartDNSServer() {
	tld := fmt.Sprintf("%s.", "test")
	log.Printf("DNS server tld: %s", tld)
	dns.HandleFunc(tld, handleRequest)
	server := &dns.Server{Addr: "127.0.0.1:53", Net: "udp"}
	log.Println("Starting DNS server on :53")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start DNS server: %v", err)
	}
}

func handleRequest(w dns.ResponseWriter, r *dns.Msg) {
	msg := dns.Msg{}
	msg.SetReply(r)

	log.Printf("Received DNS request: %v", r.Question)
	for _, q := range r.Question {
		fmt.Printf("Processing question: %s\n", q.Name)
		domain := strings.TrimSuffix(q.Name, fmt.Sprintf(".%s.", config.DefaultConfig.TLD))
		if mappingservice.FindByDomain(domain) != nil {
			log.Printf("Found mapping for: %s", q.Name)
			rr, _ := dns.NewRR(q.Name + " A 127.0.0.1")
			msg.Answer = append(msg.Answer, rr)
		}
	}

	w.WriteMsg(&msg)
}
