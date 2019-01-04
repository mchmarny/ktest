package handlers

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	// header
	w.Header().Set("Service", "tellmeall")
	w.WriteHeader(200)

	// content
	fmt.Fprintf(w, "OK")
}
