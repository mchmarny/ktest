package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func EnvVarHandler(w http.ResponseWriter, r *http.Request) {

	var request []string

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		request = append(request, fmt.Sprintf("%v: %v", pair[0], pair[1]))
	}

	fmt.Fprintf(w, strings.Join(request, "\n"))
}
