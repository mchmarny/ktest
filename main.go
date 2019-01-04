package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mchmarny/tellmeall/handlers"
)

func main() {

	// routes
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/env", handlers.EnvVarHandler)
	http.HandleFunc("/head", handlers.HeaderHandler)
	http.HandleFunc("/mem", handlers.MemoryHandler)
	http.HandleFunc("/kn", handlers.KnativeHandler)
	http.HandleFunc("/host", handlers.HostHandler)
	http.HandleFunc("/log", handlers.LogHandler)
	http.HandleFunc("/help", handlers.HelpHandler)

	// server
	port := handlers.GetPort()
	log.Printf("Starting server on %s port\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
