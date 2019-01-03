package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

func HeaderHandler(w http.ResponseWriter, r *http.Request) {

	var request []string

	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	fmt.Fprintf(w, strings.Join(request, "\n"))
}
