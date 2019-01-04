package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/http/httputil"
)

var (
	templates *template.Template
)

// InitHandlers initializes all handlers
func InitHandlers(mux *http.ServeMux) {

	// templates
	tmpls, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Error while parsing templates: %v", err)
	}
	templates = tmpls

	// static
	mux.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("static"))))

	// routes
	mux.HandleFunc("/", withLog(homeHandler))
	mux.HandleFunc("/env", withLog(envVarHandler))
	mux.HandleFunc("/head", withLog(headerHandler))
	mux.HandleFunc("/mem", withLog(memoryHandler))
	mux.HandleFunc("/kn", withLog(knativeHandler))
	mux.HandleFunc("/host", withLog(hostHandler))
	mux.HandleFunc("/log", withLog(logHandler))

	// health (Istio and other)
	mux.HandleFunc("/_healthz", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "ok")
	})
	mux.HandleFunc("/_health", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "ok")
	})

}

// withLog is a simple midleware to dump each request into log
func withLog(next http.HandlerFunc) http.HandlerFunc {
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
