package handlers

import (
	"log"
	"net/http"
	"time"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"name": "tellmeall",
		"on":   time.Now().String(),
	}

	if err := templates.ExecuteTemplate(w, "home", data); err != nil {
		log.Printf("Error in home template: %s", err)
	}

}
