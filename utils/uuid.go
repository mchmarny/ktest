package utils

import (
	"log"

	"github.com/google/uuid"
)

// NewID returns UUID v4
func NewID() string {
	id, err := uuid.NewRandom()
	if err != nil {
		log.Fatalf("Error while getting id: %v\n", err)
	}
	return id.String()
}
