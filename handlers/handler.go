package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/mchmarny/tellmeall/types"
	"github.com/mchmarny/tellmeall/utils"
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
	mux.HandleFunc("/req", withLog(requestHandler))
	mux.HandleFunc("/res", withLog(resourceHandler))
	mux.HandleFunc("/kn", withLog(knativeHandler))
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
