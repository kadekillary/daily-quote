package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// checkErr just helps to reduce boilerplate
func checkErr(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

// Quote contanis merely quote and author data
// from the response `additional formats are available from API`
type Quote struct {
	Quote  string `json:"q"`
	Author string `json:"a"`
}

// cleanQuote gets rid of beginning and ending brackets on response
func cleanQuote(q string) string {
	removeStartBracket := strings.Replace(q, "[ ", "", 1)
	removeEndBracket := strings.Replace(removeStartBracket, " ]", "", 1)
	return removeEndBracket
}

// getQuote grabs the daily quote from zenquotes.io and
// returns a slice of bytes
func getQuote() []uint8 {
	resp, err := http.Get("https://zenquotes.io/api/today")
	checkErr(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	checkErr(err)

	return body
}

// extractQuote takes a slice of bytes and returns a Quote struct
func extractQuote(body []uint8) Quote {
	var quote Quote

	data := cleanQuote(string(body))
	json.Unmarshal([]byte(data), &quote)

	return quote
}

func main() {
	quote := extractQuote(getQuote())
	// construct quote format
	formattedQuote := fmt.Sprintf("\"%s\" - %s", quote.Quote, quote.Author)
	fmt.Println(formattedQuote)
}
