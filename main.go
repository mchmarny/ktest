package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/mchmarny/tellmeall/handlers"
)

const (
	logFilePath = "/var/log/tellmeall.log"
)

func main() {

	// log to file only if the LOG_TO_FILE var is set
	if os.Getenv("LOG_TO_FILE") != "" {
		logFile, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Error while opening log file: %s - %v", logFilePath, err)
		}
		defer logFile.Close()
		log.SetOutput(logFile)
		log.Printf("Logging to file: %s", logFilePath)
	}

	// routes
	http.HandleFunc("/", withLogging(handlers.HomeHandler))
	http.HandleFunc("/env", withLogging(handlers.EnvVarHandler))
	http.HandleFunc("/head", withLogging(handlers.HeaderHandler))
	http.HandleFunc("/mem", withLogging(handlers.MemoryHandler))
	http.HandleFunc("/kn", withLogging(handlers.KnativeHandler))
	http.HandleFunc("/host", withLogging(handlers.HostHandler))
	http.HandleFunc("/log", withLogging(handlers.LogHandler))
	http.HandleFunc("/help", withLogging(handlers.HelpHandler))

	// server
	port := handlers.GetPort()
	log.Printf("Starting server on %s port\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

// log each request dump
func withLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		reqDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Println(err)
		} else {
			log.Println(string(reqDump))
		}

		next.ServeHTTP(w, r)
	}
}
