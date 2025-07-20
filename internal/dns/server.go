package dns

import (
	"log"
	"strings"

	"github.com/miekg/dns"
)

func StartDNSServer() {
	dns.HandleFunc(".", handleRequest)

	server := &dns.Server{Addr: ":53", Net: "udp"}
	log.Println("Starting DNS server on :53")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start DNS server: %v", err)
	}
}

func handleRequest(w dns.ResponseWriter, r *dns.Msg) {
	msg := dns.Msg{}
	msg.SetReply(r)

	for _, q := range r.Question {
		if strings.HasSuffix(q.Name, ".test.") {
			rr, _ := dns.NewRR(q.Name + " A 127.0.0.1")
			msg.Answer = append(msg.Answer, rr)
		}
	}

	w.WriteMsg(&msg)
}
