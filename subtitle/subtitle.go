package subtitle

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"subtitleFinder/utils"

	"github.com/gocolly/colly"
	"github.com/melbahja/got"
)

type LogFunc func(text string)

var movieDir string
var movieSubtitleName string
var movieFileName string
var index string

func GetSubtitles(
	moviePath string,
	movieUrl string,
	log LogFunc,
	addCheckButtons func(count int),
	addMoveAll func(),
) {
	moviePathSplit := strings.Split(moviePath, "/")
	movieFileName = moviePathSplit[len(moviePathSplit)-1]
	movieDir = strings.Replace(moviePath, "/"+movieFileName, "", 1)
	// movieDir = strings.Replace(moviePath, movieFileName, "", 1)
	fmt.Println("movieFileName:", movieFileName)
	fmt.Println("movieDir:", movieDir)
	movieFileName = strings.Replace(movieFileName, ".mp4", "", 1)
	movieFileName = strings.Replace(movieFileName, ".mkv", "", 1)
	movieUrlSplit := strings.Split(movieUrl, "/")
	movieSubtitleName = movieUrlSplit[len(movieUrlSplit)-1]

	log("start process")
	createDirs()
	log("subtitle directories created!")
	getSubUrls(log, addCheckButtons, addMoveAll)
}

func createDirs() {
	utils.CreateDir(movieDir + "/subs")
	utils.CreateDir(movieDir + "/subtitles")
}

func getSubUrls(log LogFunc, addCheckButtons func(count int), addMoveAll func()) {
	var urls []string
	url := "https://subscene.com/subtitles/" + movieSubtitleName
	counter := 0
	c := colly.NewCollector()
	c.OnHTML("table", func(e *colly.HTMLElement) {
		log("subtitle files scrapped!")
		e.ForEach("a", func(_ int, elem *colly.HTMLElement) {
			href := elem.Attr("href")
			isExist, _ := utils.InArray(href, urls)
			if strings.Contains(href, "farsi") && !isExist {
				urls = append(urls, href)
				getFile("https://subscene.com"+href, counter, log)
				counter++
			}
		})
		log("end of process!")
		addCheckButtons(counter)
		addMoveAll()
	})
	_ = c.Visit(url)
}

func getFile(url string, counter int, log LogFunc) {
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
		fmt.Println("@@@" + dest)
		log(dest)
		g := got.New()
		_ = g.Download(href, dest)
		utils.Unzip(dest, fld2)
	})
	_ = c.Visit(url)
}

func findSrtFile() string {
	var srtFile string
	subtleDir := movieDir + "/subtitles/" + index
	files := utils.GetDirFiles(subtleDir)
	for _, file := range files {
		if strings.Contains(file, "srt") {
			srtFile = file
		}
	}
	return srtFile
}

func moveSrtFile(srtFile string) {
	destFile := movieDir + "/" + movieFileName + ".fa.srt"
	srtBytes, _ := ioutil.ReadFile(srtFile)
	_, srtEncoding, _ := utils.DetectEncoding(srtBytes)
	fmt.Println(srtEncoding)
	if !strings.Contains(srtEncoding, "utf") {
		utils.EncodeToUTF8(srtFile, destFile)
	} else {
		utils.MoveFile(srtFile, destFile)
	}
}

func CheckSub(selectedIndex string, log LogFunc) {
	index = selectedIndex
	srtFile := findSrtFile()
	moveSrtFile(srtFile)
	log("Subtitle " + index + " was checked!")
}

func MoveAll(log LogFunc) {
	subtitlesDir := movieDir + "/subtitles/"
	srtFiles := []string{}
	err := filepath.Walk(subtitlesDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			file := path
			if strings.Contains(file, "srt") {
				srtFiles = append(srtFiles, file)
			}
			return nil
		})
	if err != nil {
		fmt.Println(err)
	}
	for i, file := range srtFiles {
		destFile := movieDir + "/" + strconv.Itoa(i) + ".fa.srt"
		fmt.Println(file, destFile)
		utils.MoveFile(file, destFile)
	}
}
