package cmd

import (
	"fmt"
	"os"
	"strings"
	"toolkit/runner/encodeRunner"
	"toolkit/utils/stdinUtil"

	"github.com/spf13/cobra"
)

type EncodeCmdData struct {
	useBase64 bool
	useURL    bool
}

var encodeCmdData EncodeCmdData

var encodeCmd = &cobra.Command{
	Use:   "encode TEXT",
	Short: "Encode a string.",
	Long: `Encode a string using Base64 or URL encoding.

Pass a string as an argument or pipe input from stdin:
  -b / --base64  Base64 encode
  -u / --url     URL (percent) encode

Examples:
  toolkit encode -b "hello world"
  toolkit encode -u "hello world"
  echo hello | toolkit encode -b`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 && !stdinUtil.IsPiped() {
			fmt.Println("Please provide a string to encode or pipe input.")
			os.Exit(2)
		}
		if !encodeCmdData.useBase64 && !encodeCmdData.useURL {
			fmt.Println("Please specify an encoding: -b (base64) or -u (url)")
			os.Exit(2)
		}
		var text string
		if len(args) == 1 {
			text = args[0]
		} else {
			data, err := stdinUtil.ReadAll()
			if err != nil {
				fmt.Printf("error reading stdin: %v\n", err)
				os.Exit(1)
			}
			text = strings.TrimRight(string(data), "\r\n")
		}
		runEncode(text)
	},
}

func init() {
	rootCmd.AddCommand(encodeCmd)
	encodeCmd.Flags().BoolVarP(&encodeCmdData.useBase64, "base64", "b", false, "Base64 encode")
	encodeCmd.Flags().BoolVarP(&encodeCmdData.useURL, "url", "u", false, "URL encode")
}

func runEncode(text string) {
	if encodeCmdData.useBase64 {
		fmt.Println(encodeRunner.Base64Encode(text))
	} else if encodeCmdData.useURL {
		fmt.Println(encodeRunner.URLEncode(text))
	}
}
