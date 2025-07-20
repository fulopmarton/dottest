package proxy

import (
	"dottest/config"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var routes = config.Mappings

func StartReverseProxy() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request for: %s\n", r.Host)
		targetPort, ok := routes[r.Host]
		target := fmt.Sprintf("http://localhost:%d", targetPort)
		if !ok {
			http.NotFound(w, r)
			return
		}
		url, _ := url.Parse(target)
		proxy := httputil.NewSingleHostReverseProxy(url)
		proxy.ServeHTTP(w, r)
	})

	log.Println("Starting reverse proxy on :80")
	err := http.ListenAndServe(":80", handler)
	if err != nil {
		log.Fatalf("Failed to start reverse proxy: %v", err)
	}
}
