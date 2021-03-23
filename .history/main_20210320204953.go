package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func getSubUrls() {
	c := colly.NewCollector()
	var urls []string
	c.OnHTML("table", func(e *colly.HTMLElement) {
		e.ForEach("a", func(_ int, elem *colly.HTMLElement) {
			href := elem.Attr("href")
			if strings.Contains(href, "farsi") {
				urls = append(urls, href)
			}
		})
		getFilesUrls(urls)
	})
	c.Visit("https://subscene.com/subtitles/camille-claudel-1915")
}

func getFilesUrls(urls []string) {
	fmt.Println(urls)
}

func main() {
	getSubUrls()
}
