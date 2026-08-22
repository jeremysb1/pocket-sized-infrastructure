package main

import (
	"reflect"
	"testing"
)

var (
	handMaidsTale = Book {Author: "Margaret Atwood", Title: "The Handmaid's Tale",}
	oryxAndCrake = Book {Author: "Margaret Atwood", Title: "Oryx and Crake"}
	theBellJar = Book {Author: "Sylvia Plath", Title: "The Bell Jar"}
	janeEyre = Book {Author: "Charlotte Brontë", Title: "Jane Eyre"}
)

func TestLoadBookworms_Success(t *testing.T) {
	type testCase struct {
		bookwormsFile string
		want          []Bookworm
		wantErr       bool
	}
}