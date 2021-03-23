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
var movieSubtitleName = "camille-claudel-1915"

func main() {
	createDirs()
	getSubUrls()
}

func createDirs() {
	utils.CreateDir(movieDir + "/subs")
	utils.CreateDir(movieDir + "subtitles")
}

func getSubUrls() {
	url := "https://subscene.com/subtitles/" + movieSubtitleName
	counter := 0
	c := colly.NewCollector()
	c.OnHTML("table", func(e *colly.HTMLElement) {
		e.ForEach("a", func(_ int, elem *colly.HTMLElement) {
			href := elem.Attr("href")
			if strings.Contains(href, "farsi") {
				getFile("https://subscene.com"+href, counter)
				counter++
			}
		})

	})
	_ = c.Visit(url)
}

func getFile(url string, counter int) {
	dir := movieDir + "/subs"
	c := colly.NewCollector()
	c.OnHTML("#downloadButton", func(e *colly.HTMLElement) {
		href := "https://subscene.com" + e.Attr("href")
		fld := dir + fmt.Sprint(counter)
		utils.CreateDir(fld)
		dest := fld + "/sub.zip"
		fmt.Println(dest)
		g := got.New()
		_ = g.Download(href, dest)
		utils.Unzip(dest, dir)
	})
	_ = c.Visit(url)
}
