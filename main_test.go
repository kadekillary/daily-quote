package main

import (
	"fmt"
	"net/http"
	"testing"
)

var casesQuote = []struct {
	In  string
	Out string
}{
	{In: "[ aksdjlfsd ]", Out: "aksdjlfsd"},
	{In: "[ jp93ji[] ]", Out: "jp93ji[]"},
	{In: "[ {'q': 'hello testing'} ]", Out: "{'q': 'hello testing'}"},
	{In: "[  ]", Out: ""},
}

func TestCleanQuote(t *testing.T) {
	for _, test := range casesQuote {
		t.Run(fmt.Sprintf("%s gets converted to %s", test.In, test.Out), func(t *testing.T) {
			got := cleanQuote(test.In)
			if got != test.Out {
				t.Errorf("got %s, want %s", got, test.Out)
			}
		})
	}
}

func TestExtractQuote(t *testing.T) {
	t.Run("check field extraction", func(t *testing.T) {
		quote := extractQuote(getQuote())

		if quote.Quote == "" {
			t.Error("Quote is missing")
		}

		if quote.Author == "" {
			t.Error("Author is missing")
		}
	})
}

func TestGetQuote(t *testing.T) {
	t.Run("check status code", func(t *testing.T) {
		resp, _ := http.Get("https://zenquotes.io/api/today")
		got := resp.StatusCode
		want := 200

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("check content", func(t *testing.T) {
		body := getQuote()
		got := len(body)
		want := 100

		if got < want {
			t.Errorf("got %d, want at least len %d", got, want)
		}
	})
}
