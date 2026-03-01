package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"
	"toolkit/runner/tail"
	"toolkit/util/file"
)

type TailCmdData struct {
	isFollow bool
	lines    int
}

var tailCmdData TailCmdData

var tailCmd = &cobra.Command{
	Use:   "tail",
	Short: "Tail a file.",
	Long: `Print the last portion of a file.

  -f / --follow  Keep watching the file and print new lines as they appear

Examples:
  toolkit tail ./app.log
  toolkit tail -f ./app.log`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Please specify a file to tail")
			os.Exit(1)
		}
		runTail(args)
	},
}

func init() {
	rootCmd.AddCommand(tailCmd)

	tailCmd.Flags().BoolVarP(&tailCmdData.isFollow, "follow", "f", false, "Continue looking for new lines")
	tailCmd.Flags().IntVarP(&tailCmdData.lines, "lines", "n", 10, "Number of lines to show")
}

func runTail(args []string) {
	logFile := args[0]

	lastPosition, err := tail.SeekLines(logFile, tailCmdData.lines)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	currentPosition, err := file.GetFileSize(logFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	content, lastPosition := tail.ReadFile(logFile, lastPosition, currentPosition)
	fmt.Print(content)

	if !tailCmdData.isFollow {
		return
	}

	c := time.Tick(100 * time.Millisecond)
	for range c {
		newPosition, err := file.GetFileSize(logFile)
		if err == nil {
			currentPosition = newPosition
		}

		if lastPosition == currentPosition {
			continue
		}
		content, lastPosition = tail.ReadFile(logFile, lastPosition, currentPosition)
		fmt.Print(content)
	}
}
