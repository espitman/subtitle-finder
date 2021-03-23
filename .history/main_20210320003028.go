package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func Contains(s, substr string) bool {
	return Index(s, substr) >= 0
}

func main() {
	c := colly.NewCollector()
	// Find and visit all links
	c.OnHTML("tr td.a1 a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://subscene.com/subtitles/camille-claudel-1915")
}
