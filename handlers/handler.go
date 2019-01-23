package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/mchmarny/ktest/types"
	"github.com/mchmarny/ktest/utils"
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
	mux.HandleFunc("/", withRequestLog(homeHandler))
	mux.HandleFunc("/req", withRequestLog(requestHandler))
	mux.HandleFunc("/res", withRequestLog(resourceHandler))
	mux.HandleFunc("/kn", withRequestLog(knativeHandler))
	mux.HandleFunc("/log", withRequestLog(logHandler))

	// health (Istio and other)
	mux.HandleFunc("/_healthz", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "ok")
	})

}

// withRequestLog is a simple midleware to dump each request into log
func withRequestLog(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		notCache(w)

		reqDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Println(err)
		} else {
			log.Println(string(reqDump))
		}

		next.ServeHTTP(w, r)
	}
}

func getMeta(r *http.Request) *types.RequestMetadata {
	return &types.RequestMetadata{
		ID:     utils.NewID(),
		Ts:     time.Now(),
		URI:    r.RequestURI,
		Host:   r.Host,
		Method: r.Method,
	}
}

// WriteObject write content to response
func writeJSON(w http.ResponseWriter, o interface{}) {

	w.WriteHeader(http.StatusOK)
	e := json.NewEncoder(w)
	e.SetEscapeHTML(true)
	e.SetIndent("", "\t")
	e.Encode(o)

}

func notCache(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
}
