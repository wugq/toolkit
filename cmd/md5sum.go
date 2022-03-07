package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"toolkit/runner/md5sumRunner"
	"toolkit/utils/fileUtil"
)

type Md5sumCmdData struct {
	text string
}

var md5sumCmdData Md5sumCmdData
var md5sumCmd = &cobra.Command{
	Use:   "md5sum FILE",
	Short: "check md5sum of a file or string.",
	Long:  `check md5sum of a file or string.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Println("Please process one file at a time.")
			os.Exit(1)
		}
		runMd5Sum(args)
	},
}

func init() {
	rootCmd.AddCommand(md5sumCmd)

	md5sumCmd.Flags().StringVarP(&md5sumCmdData.text, "text", "t", "", "specify a Text string")
}

func runMd5Sum(args []string) {
	if len(args) == 1 {
		userInput := args[0]
		isFile, err := fileUtil.IsFile(userInput)
		if !isFile || err != nil {
			fmt.Println("File is incorrect.")
		}

		checksum, err := md5sumRunner.CheckFile(userInput)
		if err != nil {
			fmt.Printf("error: %v", err)
			return
		}
		fmt.Printf("%v %v", checksum, userInput)
	} else if md5sumCmdData.text != "" {
		checksum, err := md5sumRunner.CheckText(md5sumCmdData.text)
		if err != nil {
			fmt.Printf("error: %v", err)
			return
		}
		fmt.Printf("%v %v", checksum, md5sumCmdData.text)
	} else {
		fmt.Println("Nothing to do.")
	}
}
