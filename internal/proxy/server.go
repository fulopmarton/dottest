package proxy

import (
	"crypto/tls"
	mappingservice "dottest/internal/services"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"runtime/debug"
	"strings"
)

// TODO: Implement a thread safe cache for certifications, add it to loadCertForDomain
func loadCertForDomain(domain string) (*tls.Certificate, error) {
	certFile := fmt.Sprintf("./data/certs/%s.pem", domain)
	keyFile := fmt.Sprintf("./data/certs/%s-key.pem", domain)

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}
	return &cert, nil
}

func getCertificate(clientHello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	cert, err := loadCertForDomain(clientHello.ServerName)
	if err != nil {
		log.Printf("Failed to load certificate for domain %s: %v", clientHello.ServerName, err)
		return nil, err
	}
	return cert, nil
}

func StartReverseProxy() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Panic recovered: %v\n%s", r, debug.Stack())
		}
	}()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		domain := strings.TrimSuffix(r.Host, ".test")
		mapping := mappingservice.FindByDomain(domain)
		if mapping == nil {
			http.NotFound(w, r)
			return
		}
		url, _ := url.Parse(mapping.Target)
		proxy := httputil.NewSingleHostReverseProxy(url)
		proxy.ServeHTTP(w, r)
	})

	log.Println("Starting reverse proxy on :80")
	go func() {
		err := http.ListenAndServe(":80", handler)
		if err != nil {
			fmt.Printf("Failed to start HTTP server: %v\n", err)
		}
	}()

	tlsConfig := &tls.Config{
		GetCertificate: getCertificate,
		MinVersion:     tls.VersionTLS12,
	}
	server := &http.Server{
		Addr:      ":443",
		Handler:   handler,
		TLSConfig: tlsConfig,
	}
	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", server.Addr, err)
	}
	tlsListener := tls.NewListener(ln, tlsConfig)
	log.Printf("Starting reverse proxy on %s", server.Addr)
	log.Fatal(server.Serve(tlsListener))
}
