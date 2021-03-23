package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
	"github.com/melbahja/got"
	"github.com/spf13/cobra"
)

var movieDir = "/Users/osx/Desktop/a"
var movieFileName = "camille-claudel-1915-YIFY.mp4"
var movieSubtitleName = "camille-claudel-1915"

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

var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "about go gin starter kit",
	Long:  `about go gin starter kit`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("subtitle finder created by saeed heidari")
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}
