package main

import (
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
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := loadBookworms(tc.bookwormsFile)

			if tc.wantErr {
				if err == nil {
					t.Fatal("expected err, got nothing")
				}
				return
			}

			// we aren't expecting errors here, this should be the happy path
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if !equalBookworms(t, got, tc.want) {
				t.Fatalf("different result: got %v, expected %v", got, tc.want)
			}
		})
	}
}

// 