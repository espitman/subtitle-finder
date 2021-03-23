package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func getSubUrls() {
	c := colly.NewCollector()
	c.OnHTML("table", func(e *colly.HTMLElement) {
		e.ForEach("a", func(_ int, elem *colly.HTMLElement) {
			href := elem.Attr("href")
			if strings.Contains(href, "farsi") {
				getFilesUrl(href)
			}
		})

	})
	c.Visit("https://subscene.com/subtitles/camille-claudel-1915")
}

func getFilesUrl(url string) {
	c := colly.NewCollector()
	c.OnHTML("#downloadButton", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		fmt.Println(href)
	})
	_ = c.Visit(url)
}

func main() {
	getSubUrls()
}
