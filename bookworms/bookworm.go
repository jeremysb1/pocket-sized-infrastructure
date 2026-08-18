package main

import (
	"encoding/json"
	"os"
)

// Bookworm contains the list of books on a bookworm's shelf.
type Bookworm struct {
	Name string `json:"name"`
	Books []Book `json:"books"`
}

// Book describes a book on a bookworm's shelf
type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

// loadBookworms reads the file and returns the list of
// bookworms and their books
func loadBookworms(filePath string) ([]Bookworm, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
}
