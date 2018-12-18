package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	defaultPort = "8080"
)

func handleDefault(w http.ResponseWriter, r *http.Request) {

	var request []string

	// method, url, proto
	request = append(request, "METHOD/URL/PROTO")
	request = append(request, "================")
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	request = append(request, fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto))
	request = append(request, "\n")

	// headers
	request = append(request, "HEADER")
	request = append(request, "======")
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}
	request = append(request, "\n")

	// env vars
	request = append(request, "ENV VARS")
	request = append(request, "========")
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		request = append(request, fmt.Sprintf("%v: %v", pair[0], pair[1]))
	}
	request = append(request, "\n")

	fmt.Fprintf(w, strings.Join(request, "\n"))
}

func main() {
	http.HandleFunc("/", handleDefault)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	log.Printf("Starting server on %s port\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
