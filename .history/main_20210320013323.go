package main

import (
	"strings"

	"github.com/gocolly/colly"
)

func getSubUrls() {
	c := colly.NewCollector()
	var urls []string
	c.OnHTML("table", func(e *colly.HTMLElement) {
		e.ForEach("a", func(_ int, elem *colly.HTMLElement) {
			if strings.Contains(elem.Attr("href"), "farsi") {
				urls = append(urls, elem.Attr("href"))
			}
		})
	})

	c.Visit("https://subscene.com/subtitles/camille-claudel-1915")
}

func main() {
	getSubUrls()
}
