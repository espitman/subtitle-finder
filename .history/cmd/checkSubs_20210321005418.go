package cmd

import (
	"fmt"
	"os"

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

		fmt.Println("check", index)
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(checkSubCmd)
}