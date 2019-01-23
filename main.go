package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mchmarny/ktest/handlers"
	"github.com/mchmarny/ktest/utils"
)

func main() {

	// log
	utils.ConfigureLogging()
	defer utils.FinalizeLogging()

	// mux
	mux := http.NewServeMux()

	// routes
	handlers.InitHandlers(mux)

	// port
	port := utils.MustGetEnv("PORT", "8080")

	// server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}

	log.Printf("Server starting on port %s \n", port)
	log.Fatal(server.ListenAndServe())
}


