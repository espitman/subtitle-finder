package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"subtitleFinder/utils"

	"github.com/spf13/cobra"
)

var checkSubCmd = &cobra.Command{
	Use:   "check",
	Short: "about go gin starter kit",
	Long:  `about go gin starter kit`,
	Run: func(cmd *cobra.Command, args []string) {

		movieDir = args[0]
		movieFileName = args[1]
		index = args[2]

		checkSub()
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(checkSubCmd)
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
	utils.MoveFile(srtFile, destFile)
}

func checkSub() {
	srtFile := findSrtFile()
	fmt.Println(srtFile)
	a, _ := ioutil.ReadFile(srtFile)
	b, c, _ := utils.DetectEncoding(a, "")
	fmt.Println(b, c)
	// moveSrtFile(srtFile)
}
