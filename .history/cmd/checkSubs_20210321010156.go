package cmd

import (
	"os"
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

func checkSub() {
	subtleDir := movieDir + "/subtitles/" + index
	utils.GetDirFiles(subtleDir)
}
