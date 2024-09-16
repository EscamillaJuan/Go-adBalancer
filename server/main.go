package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func handler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello from server on port %s\n", req.Host)
}

func main() {
	port := flag.String("port", "8080", "port to listen on")
	flag.Parse()

	http.HandleFunc("/", handler)

	log.Printf("Server listening on port %s", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}