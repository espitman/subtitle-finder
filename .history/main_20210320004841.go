package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func getSubUrls() {
	c := colly.NewCollector()
	var urls []string
	// Find and visit all links
	c.OnHTML("tr td.a1 a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		url := r.URL.String()
		if strings.Contains(url, "farsi") {
			urls = append(urls, url)
		}
		fmt.Println("Visiting", urls)
	})

	c.
		c.Visit("https://subscene.com/subtitles/camille-claudel-1915")
}

func main() {
	getSubUrls()
}
