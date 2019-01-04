package utils

import (
	"testing"
)

func TestPing(t *testing.T) {
	testID := NewID()
	if testID == "" || len(testID) < 10 {
		t.Fatal("Invalid ID generated")
	}
}
