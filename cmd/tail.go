package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"
	"toolkit/runner/tailRunner"
	"toolkit/utils/fileUtil"
)

type TailCmdData struct {
	isFollow bool
}

var tailCmdData TailCmdData

var tailCmd = &cobra.Command{
	Use:   "tail",
	Short: "Tail a file.",
	Long:  `Tail a file.`,
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
}

func runTail(args []string) {
	logFile := args[0]
	const buffSize = 100

	currentPosition, err := fileUtil.GetFileSize(logFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lastPosition := currentPosition - buffSize

	lastPosition = tailRunner.ReadFile(logFile, lastPosition, currentPosition)

	if !tailCmdData.isFollow {
		fmt.Println("!isFollow quit")
		return
	}

	c := time.Tick(100 * time.Millisecond)
	for _ = range c {
		newPosition, err := fileUtil.GetFileSize(logFile)
		if err == nil {
			currentPosition = newPosition
		}

		if lastPosition == currentPosition {
			continue
		}
		lastPosition = tailRunner.ReadFile(logFile, lastPosition, currentPosition)
	}
}
