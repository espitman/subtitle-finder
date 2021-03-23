package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	"github.com/melbahja/got"
)

func getSubUrls() {
	c := colly.NewCollector()
	c.OnHTML("table", func(e *colly.HTMLElement) {
		e.ForEach("a", func(_ int, elem *colly.HTMLElement) {
			href := elem.Attr("href")
			if strings.Contains(href, "farsi") {
				getFilesUrl("https://subscene.com" + href)
			}
		})

	})
	_ = c.Visit("https://subscene.com/subtitles/camille-claudel-1915")
}

func getFilesUrl(url string) {
	dir := "/Users/osx/Desktop/a/"
	// fmt.Println("::", url)
	c := colly.NewCollector()
	c.OnHTML("#downloadButton", func(e *colly.HTMLElement) {
		href := "https://subscene.com" + e.Attr("href")
		fmt.Println(href)
		// _ = utils.DownloadFile(href, dir+"1.zip")
		g := got.New()
		err := g.Download(href, dir+"1.zip")
		fmt.Println(err)
	})
	_ = c.Visit(url)
}

func main() {
	getSubUrls()
}