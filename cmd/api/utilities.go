package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func fetchDescription(name string) string {

	var result string
	c := colly.NewCollector()
	c.OnHTML(".mw-parser-output", func(e *colly.HTMLElement) {

		found := false
		e.ForEach("p", func(_ int, elem *colly.HTMLElement) {
			if !found && strings.Contains(elem.Text, name) {
				result = elem.Text
				found = true
			}
		})
	})
	fmt.Println("Scrapping Complete")

	c.Visit("https://en.wikipedia.org/wiki/" + name)
	return result
}
