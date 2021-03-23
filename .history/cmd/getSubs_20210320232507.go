package cmd

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func createDirs() {
	utils.CreateDir(movieDir + "/subs")
	utils.CreateDir(movieDir + "/subtitles")
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
	subsDir := movieDir + "/subs/"
	subtitlesDir := movieDir + "/subtitles/"
	c := colly.NewCollector()
	c.OnHTML("#downloadButton", func(e *colly.HTMLElement) {
		href := "https://subscene.com" + e.Attr("href")
		fld := subsDir + fmt.Sprint(counter)
		fld2 := subtitlesDir + fmt.Sprint(counter)
		utils.CreateDir(fld)
		utils.CreateDir(fld2)
		dest := fld + "/sub.zip"
		fmt.Println(dest)
		g := got.New()
		_ = g.Download(href, dest)
		utils.Unzip(dest, fld2)
	})
	_ = c.Visit(url)
}