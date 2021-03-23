package cmd

import (
	"fmt"
	"os"
	"strings"

	"subtitleFinder/utils"

	"github.com/gocolly/colly"
	"github.com/melbahja/got"
	"github.com/spf13/cobra"
)

var movieDir string
var movieSubtitleName string
var movieFileName string
var index string

var getSubCmd = &cobra.Command{
	Use:   "get",
	Short: "about go gin starter kit",
	Long:  `about go gin starter kit`,
	Run: func(cmd *cobra.Command, args []string) {

		movieDir = args[0]
		movieSubtitleName = args[1]

		fmt.Println("get")
		createDirs()
		getSubUrls()
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(getSubCmd)
}

func createDirs() {
	utils.CreateDir(movieDir + "/subs")
	utils.CreateDir(movieDir + "/subtitles")
}

func getSubUrls() {
	var urls []string
	url := "https://subscene.com/subtitles/" + movieSubtitleName
	counter := 0
	c := colly.NewCollector()
	c.OnHTML("table", func(e *colly.HTMLElement) {
		e.ForEach("a", func(_ int, elem *colly.HTMLElement) {
			href := elem.Attr("href")
			if strings.Contains(href, "farsi" && !utils.InArray(href, urls)) {
				urls = append(urls, href)
				fmt.Println(href)
				// getFile("https://subscene.com"+href, counter)
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
