package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

type ServerPool struct {
	servers []*url.URL
	current uint64
}

func (p *ServerPool) getNextServer() *url.URL {
	next := atomic.AddUint64(&p.current, 1)
	return p.servers[next % uint64(len(p.servers))]
}

func (p *ServerPool) proxyHandler(res http.ResponseWriter, req *http.Request) {
	target := p.getNextServer()
	log.Printf("Redirect to: %s", target)

	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(res, req)
}

func main() {
	servers := []string {
		"http://localhost:8081",
		"http://localhost:8082",
		"http://localhost:8083",
	}

	var serverPool ServerPool
	for _, s := range servers {
		serverUrl, err := url.Parse(s)
		if err != nil {
			log.Fatalf("Error parsing server URL %s: %v ", s, err)
		}

		serverPool.servers = append(serverPool.servers, serverUrl)
	}

	http.HandleFunc("/", serverPool.proxyHandler)
	fmt.Println("Go-ad balancer listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}