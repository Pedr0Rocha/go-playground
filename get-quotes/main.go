package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Quote struct {
	ID           string   `json:"_id"`
	Tags         []string `json:"tags"`
	Content      string   `json:"content"`
	Author       string   `json:"author"`
	AuthorSlug   string   `json:"authorSlug"`
	Length       int      `json:"length"`
	DateAdded    string   `json:"dateAdded"`
	DateModified string   `json:"dateModified"`
}

func main() {
	quote, err := GetRandomQuote()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%q by %s\n", quote.Content, quote.Author)
}

func GetRandomQuote() (*Quote, error) {
	resp, err := http.Get("https://api.quotable.io/random")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var quote Quote
	if err := json.Unmarshal(body, &quote); err != nil {
		return nil, err
	}

	return &quote, nil
}
