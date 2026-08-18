package main

import (
	"encoding/json"
	"os"
)

// A Bookworm contains a list of books on a bookworm's shelf.
type Bookworm struct {
	Name string `json:"name"`
	Books []Book `json:"books"`
}

// loadBookworms reads the file and returns the list of
// bookworms and their books
func loadBookworms(filePath string) ([]Bookworm, error) {
	return nil, nil
}
