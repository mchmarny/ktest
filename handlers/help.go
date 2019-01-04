package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func HelpHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Handling Help...")

	// header
	w.Header().Set("Service", "tellmeall")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(200)

	// content
	fmt.Fprintf(w, "<h3>Valid endpoints</h3>")
	fmt.Fprintf(w, "<ul>")
	fmt.Fprintf(w, "<li><a href='/'>/</a> responds with 'OK' (ala healthcheck)</li>")
	fmt.Fprintf(w, "<li><a href='/env'>/env</a> responds with all environment variables in a key/value format</li>")
	fmt.Fprintf(w, "<li><a href='/head'>/head</a> responds with all request header variables in a key/value format</li>")
	fmt.Fprintf(w, "<li><a href='/mem'>/mem</a> responds with total, used and free system memory information</li>")
	fmt.Fprintf(w, "<li><a href='/host'>/host</a> responds with container info (ID, Hostname, OS, Boot-time etc.)</li>")
	fmt.Fprintf(w, "<li><a href='/log'>/log</a> responds with content of specific log or log dir (e.g. /log?logpath=/var/log/app.log)</li>")
	fmt.Fprintf(w, "<li><a href='/kn'>/kn</a> responds with Knative-specific data as defined in the [Runtime Contract](https://github.com/knative/serving/blob/master/docs/runtime-contract.md)</li>")
	fmt.Fprintf(w, "<li><a href='/help'>/help</a> responds with this list of endpoints as URLs</li>")
	fmt.Fprintf(w, "</ul>")
}
