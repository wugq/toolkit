package cmd

import (
	"fmt"
	"os"
	"strings"
	"toolkit/runner/encodeRunner"
	"toolkit/utils/stdinUtil"

	"github.com/spf13/cobra"
)

type DecodeCmdData struct {
	useBase64 bool
	useURL    bool
}

var decodeCmdData DecodeCmdData

var decodeCmd = &cobra.Command{
	Use:   "decode TEXT",
	Short: "Decode a string.",
	Long: `Decode a Base64 or URL-encoded string.

Pass a string as an argument or pipe input from stdin:
  -b / --base64  Base64 decode
  -u / --url     URL (percent) decode

Examples:
  toolkit decode -b "aGVsbG8gd29ybGQ="
  toolkit decode -u "hello+world"
  echo aGVsbG8= | toolkit decode -b`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 && !stdinUtil.IsPiped() {
			fmt.Println("Please provide a string to decode or pipe input.")
			os.Exit(2)
		}
		if !decodeCmdData.useBase64 && !decodeCmdData.useURL {
			fmt.Println("Please specify a decoding: -b (base64) or -u (url)")
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
		runDecode(text)
	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)
	decodeCmd.Flags().BoolVarP(&decodeCmdData.useBase64, "base64", "b", false, "Base64 decode")
	decodeCmd.Flags().BoolVarP(&decodeCmdData.useURL, "url", "u", false, "URL decode")
}

func runDecode(text string) {
	if decodeCmdData.useBase64 {
		result, err := encodeRunner.Base64Decode(text)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(result)
	} else if decodeCmdData.useURL {
		result, err := encodeRunner.URLDecode(text)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(result)
	}
}
