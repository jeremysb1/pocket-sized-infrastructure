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
	tests := map[string]testCase{
		"file exists": {
			bookwormsFile: "testdata/bookworms.json",
			want: []Bookworm{
				{Name: "Fadi", Books: []Book{handMaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handMaidsTale, janeEyre}},
			},
			wantErr: false,
		},
		"file doesn't exist": {
			bookwormsFile: "testdata/bookworms.json",
			want: nil,
			wantErr: true,
		},
		"invalid JSON": {
			bookwormsFile: "testdata/bookworms.json",
			want: nil,
			wantErr: true,
		},
	}
}