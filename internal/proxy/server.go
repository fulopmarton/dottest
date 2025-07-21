package proxy

import (
	mappingservice "dottest/internal/services"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func StartReverseProxy() {
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
	err := http.ListenAndServe(":80", handler)
	if err != nil {
		log.Fatalf("Failed to start reverse proxy: %v", err)
	}
	// https
	certFile := fmt.Sprintf("./%s.pem")
	keyFile := fmt.Sprintf("./%s-key.pem")
	http.ListenAndServeTLS(":443", certFile, keyFile, handler)
}
