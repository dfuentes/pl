package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func main() {
	conf, err := Load("./test.txt")
	if err != nil {
		log.Fatal(err)
	}

	client := NewClient()

	comics, err := client.GetNewReleases()
	if err != nil {
		log.Fatal(err)
	}

	comics = Filter(comics, conf.Subscriptions)

	for _, c := range comics {
		fmt.Println(c.Title)
	}
}

func Filter(comics []ComicDetails, titles []string) []ComicDetails {
	pattern := "(?i)^("
	pattern += strings.Join(titles, "|")
	pattern += `)\s+\#\d+`

	regex, err := regexp.Compile(pattern)
	if err != nil {
		return []ComicDetails{}
	}

	result := []ComicDetails{}

	for _, comic := range comics {
		if regex.MatchString(comic.Title) {
			result = append(result, comic)
		}
	}

	return result
}
