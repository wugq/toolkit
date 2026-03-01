package cmd

import (
	"fmt"
	"os"
	"toolkit/runner/json"
	"toolkit/utils/stdinutil"

	"github.com/spf13/cobra"
)

type JsonCmdData struct {
	text string
}

var jsonCmdData JsonCmdData

var jsonCmd = &cobra.Command{
	Use:   "json FILE",
	Short: "Pretty-print and validate JSON.",
	Long: `Pretty-print and validate a JSON file or string.

Provide a file path, use -t for an inline string, or pipe JSON input:
  -t / --text  JSON string to pretty-print

Examples:
  toolkit json ./data.json
  toolkit json -t '{"a":1,"b":2}'
  type data.json | toolkit json`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 && jsonCmdData.text == "" && !stdinutil.IsPiped() {
			fmt.Println("Please provide a file path, use -t, or pipe input.")
			os.Exit(2)
		}
		runJson(args)
	},
}

func init() {
	rootCmd.AddCommand(jsonCmd)
	jsonCmd.Flags().StringVarP(&jsonCmdData.text, "text", "t", "", "JSON string to pretty-print")
}

func runJson(args []string) {
	var input []byte

	if jsonCmdData.text != "" {
		input = []byte(jsonCmdData.text)
	} else if stdinutil.IsPiped() {
		var err error
		input, err = stdinutil.ReadAll()
		if err != nil {
			fmt.Printf("error reading stdin: %v\n", err)
			os.Exit(1)
		}
	} else {
		var err error
		input, err = os.ReadFile(args[0])
		if err != nil {
			fmt.Printf("error reading file: %v\n", err)
			os.Exit(1)
		}
	}

	result, err := json.PrettyPrint(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result)
}
