package main

import (
	"strings"

	"github.com/gocolly/colly"
)

func getSubUrls() {
	counter := 0
	c := colly.NewCollector()
	c.OnHTML("table", func(e *colly.HTMLElement) {
		e.ForEach("a", func(_ int, elem *colly.HTMLElement) {
			href := elem.Attr("href")
			if strings.Contains(href, "farsi") {
				getFilesUrl("https://subscene.com"+href, counter)
				counter++
			}
		})

	})
	_ = c.Visit("https://subscene.com/subtitles/camille-claudel-1915")
}

func getFilesUrl(url string, counter int) {
	dir := "/Users/osx/Desktop/a/"
	// fmt.Println("::", url)
	c := colly.NewCollector()
	c.OnHTML("#downloadButton", func(e *colly.HTMLElement) {
		href := "https://subscene.com" + e.Attr("href")
		g := got.New()
		_ = g.Download(href, dir+"1.zip")
		// fmt.Println(err)
	})
	_ = c.Visit(url)
}

func main() {
	getSubUrls()
}
