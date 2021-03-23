package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func getSubUrls() {
	// var count = 0
	c := colly.NewCollector()
	var urls []string
	// Find and visit all links
	c.OnHTML("tr td.a1 a", func(e *colly.HTMLElement) {
		// count++
		// e.Request.Visit(e.Attr("href"))
		link := e.Request.AbsoluteURL(e.Attr("href"))
		if strings.Contains(link, "farsi") {
			urls = append(urls, link)
			fmt.Println("link", link)
		}
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("end")
	})

	// fmt.Println("urls", urls)

	// c.OnRequest(func(r *colly.Request) {
	// 	url := r.URL.String()
	// 	fmt.Println("url", url)
	// 	if strings.Contains(url, "farsi") {
	// 		urls = append(urls, url)
	// 	}

	// })

	// c.OnScraped(func(r *colly.Response) {
	// 	fmt.Println("Finished", r.Request.URL)
	// 	fmt.Println("Visiting", urls)
	// })

	c.Visit("https://subscene.com/subtitles/camille-claudel-1915")
}

func main() {
	getSubUrls()
}
