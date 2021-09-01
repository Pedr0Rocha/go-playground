package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	quote, err := GetRandomQuote()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(quote)
}

func GetRandomQuote() (string, error) {
	resp, err := http.Get("https://api.quotable.io/random")
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
