package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Handling Home...")

	// header
	w.Header().Set("Service", "tellmeall")
	w.WriteHeader(200)

	// content
	fmt.Fprintf(w, "OK")
}
