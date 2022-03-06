package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"toolkit/runner/md5sumRunner"
	"toolkit/utils/fileUtil"
)

var md5sumCmd = &cobra.Command{
	Use:   "md5sum FILE",
	Short: "check md5sum of a file or string.",
	Long:  `check md5sum of a file or string.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Please specify a file or a string")
			os.Exit(1)
		}
		runMd5Sum(args)
	},
}

func init() {
	rootCmd.AddCommand(md5sumCmd)
}

func runMd5Sum(args []string) {
	userInput := args[0]
	isFile, _ := fileUtil.IsFile(userInput)
	var checksum string
	var err error
	if isFile {
		checksum, err = md5sumRunner.CheckFile(userInput)
	} else {
		checksum, err = md5sumRunner.CheckText(userInput)
	}

	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Printf("%v %v", checksum, userInput)

}
