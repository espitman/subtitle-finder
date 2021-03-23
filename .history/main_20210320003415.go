package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	// Find and visit all links
	c.OnHTML("tr td.a1 a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		url := r.URL
		if strings.Contains(url, "farsi") {
			fmt.Println("Visiting", r.URL)
		}
	})

	c.Visit("https://subscene.com/subtitles/camille-claudel-1915")
}
