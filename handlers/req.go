package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/mchmarny/tellmeall/types"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Handling Request...")

	w.Header().Set("Content-Type", "application/json")

	sr := &types.SimpleRequest{
		Meta:    getMeta(r),
		Headers: make(map[string]interface{}),
		EnvVars: make(map[string]interface{}),
	}

	// env vars
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		sr.EnvVars[pair[0]] = pair[1]
	}

	// headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for i, h := range headers {
			if len(headers) > 1 {
				sr.Headers[fmt.Sprintf("%s[%d]", name, i)] = h
			} else {
				sr.Headers[fmt.Sprintf("%s", name)] = h
			}
		}
	}

	writeJSON(w, sr)

}
