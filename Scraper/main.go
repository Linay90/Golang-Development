package main

import (
	"fmt"
	"math/rand"
)

var googleDomains = map[string]string{}

type SearchResult struct {
	ResultRank  int
	ResultURL   string
	ResultTitle string
	ResultDesc  string
}

var userAgents = []string{}

func randomUserAgent() string {
	random := rand.Int() % len(userAgents)
	return userAgents[random]
}

func GoogleScrape() ([]SearchResult, err) {
	results := []SearchResults{}

}

func main() {
	res, err := GoogleScrape("anything")
	if err == nil {
		for _, res := range res {
			fmt.Println(res)
		}
	}

}
