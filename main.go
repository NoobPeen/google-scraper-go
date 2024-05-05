package main

import (
	"fmt"
	"math/rand"
)

var GoogleDomains = map[string]string{}

type SearchResult struct {
	ResultRank int
	ResultURL  string
	ResultDesc string
}

var userAgents = []string{}

func randomUserAgent() string {
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]
}

func buildGoogleUrls() {

}

func GoogleScrape() ([]SearchResult, err) {
	results := SearchResult{}
	resultCounter := 0
	buildGoogleUrls()

}

func main() {
	res, err := GoogleScrape("Jennifer Connelly")
	if err == nil {
		for _, res := range res {
			fmt.Println(res)
		}
	}

}
