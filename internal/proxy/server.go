package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var routes = map[string]string{
	"app.test": "http://localhost:3000",
	"api.test": "http://localhost:4000",
}

func StartReverseProxy() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		target, ok := routes[r.Host]
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

