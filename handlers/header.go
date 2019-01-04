package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func HeaderHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Handling Headers...")

	var request []string

	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	fmt.Fprintf(w, strings.Join(request, "\n"))
}
