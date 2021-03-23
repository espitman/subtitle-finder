package cmd

import (
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

}

func checkSub() {
	srtFile := findSrtFile()
	destFile := movieDir + "/" + movieFileName + "fa.srt"
	MoveFile(srtFile, destFile)
}
