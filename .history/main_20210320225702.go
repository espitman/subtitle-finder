package main

import (
	"fmt"
	"strings"
	"subtitleFinder/utils"

	"github.com/gocolly/colly"
	"github.com/melbahja/got"
)

var movieDir = "/Users/osx/Desktop/a"
var movieFileName = "camille-claudel-1915-YIFY.mp4"
var movieSubtitleName = ""

func main() {
	getSubUrls()
}

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
		fld := dir + fmt.Sprint(counter)
		utils.CreateDir(fld)
		dest := fld + "/sub.zip"
		fmt.Println(dest)
		g := got.New()
		_ = g.Download(href, dest)
		// fmt.Println(err)
		utils.Unzip(dest, dir)
	})
	_ = c.Visit(url)
}
