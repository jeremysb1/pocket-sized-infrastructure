package main

import "testing"

func TestGreet(t *testing.T) {
	type testCase struct {
		lang language
		want string
	}

	var tests = map[string]testCase{
		"Spanish": {
			lang: "es",
			want: "Hola mundo",
		},
		"English": {
			lang: "en",
			want: "Hello World",
		},
		"Akkadian, not supported": {
			lang: "akk",
			want: `unsupported language: "akk"`,
		},
		"French": {
			lang: "fr",
			want: "Bonjour le monde",
		},
		"Italian": {
			lang: "it",
			want: "Ciao a tutti!",
		},
		"Empty": {
			lang: "",
			want: `unsupported language: ""`,
		},
	}
	// range over all scenarios
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := greet(tc.lang)
			if got != tc.want {
				t.Errorf("expected: %q, got: %q", tc.want, got)
			}
		})
	}
}
