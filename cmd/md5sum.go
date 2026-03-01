package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"toolkit/runner/md5sumRunner"
	"toolkit/utils/fileUtil"
	"toolkit/utils/stdinUtil"
)

type Md5sumCmdData struct {
	text string
	algo string
}

var md5sumCmdData Md5sumCmdData
var md5sumCmd = &cobra.Command{
	Use:   "md5sum FILE",
	Short: "check md5sum of a file or string.",
	Long: `Compute the MD5 checksum of a file or a text string.

Provide a file path as an argument, use -t to hash a string, or pipe input:
  -t / --text    Compute the checksum of the given text string
  -a / --algo    Hash algorithm: md5 (default), sha256, sha512

Examples:
  toolkit md5sum ./file.txt
  toolkit md5sum -t "hello world"
  toolkit md5sum -a sha256 ./file.txt
  echo hello | toolkit md5sum`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Println("Please process one file at a time.")
			os.Exit(1)
		}
		if len(args) == 0 && md5sumCmdData.text == "" && !stdinUtil.IsPiped() {
			fmt.Println("Please provide a file, use -t, or pipe input.")
			os.Exit(2)
		}
		runMd5Sum(args)
	},
}

func init() {
	rootCmd.AddCommand(md5sumCmd)

	md5sumCmd.Flags().StringVarP(&md5sumCmdData.text, "text", "t", "", "specify a Text string")
	md5sumCmd.Flags().StringVarP(&md5sumCmdData.algo, "algo", "a", "md5", "hash algorithm: md5, sha256, sha512")
}

func runMd5Sum(args []string) {
	if len(args) == 1 {
		userInput := args[0]
		isFile, err := fileUtil.IsFile(userInput)
		if !isFile || err != nil {
			fmt.Println("File is incorrect.")
		}

		checksum, err := md5sumRunner.CheckFile(userInput, md5sumCmdData.algo)
		if err != nil {
			fmt.Printf("error: %v", err)
			return
		}
		fmt.Printf("%v %v", checksum, userInput)
	} else if md5sumCmdData.text != "" {
		checksum, err := md5sumRunner.CheckText(md5sumCmdData.text, md5sumCmdData.algo)
		if err != nil {
			fmt.Printf("error: %v", err)
			return
		}
		fmt.Printf("%v %v", checksum, md5sumCmdData.text)
	} else if stdinUtil.IsPiped() {
		data, err := stdinUtil.ReadAll()
		if err != nil {
			fmt.Printf("error reading stdin: %v\n", err)
			os.Exit(1)
		}
		checksum, err := md5sumRunner.CheckText(string(data), md5sumCmdData.algo)
		if err != nil {
			fmt.Printf("error: %v", err)
			return
		}
		fmt.Printf("%v -", checksum)
	}
}
