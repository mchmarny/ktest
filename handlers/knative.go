package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

const (
	defaultPort = "8080"
)

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	return port
}

func KnativeHandler(w http.ResponseWriter, r *http.Request) {

	var request []string

	request = append(request, fmt.Sprintf("PORT: %v", GetPort()))
	request = append(request, fmt.Sprintf("K_SERVICE: %v", os.Getenv("K_SERVICE")))
	request = append(request, fmt.Sprintf("K_REVISION: %v", os.Getenv("K_REVISION")))
	request = append(request, fmt.Sprintf("K_CONFIGURATION: %v", os.Getenv("K_CONFIGURATION")))

	fmt.Fprintf(w, strings.Join(request, "\n"))
}
