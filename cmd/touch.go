package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"
	"toolkit/runner/touchRunner"
	"toolkit/utils/fileUtil"
)

type TouchCmdData struct {
	isRecursive bool
}

var touchCmdData TouchCmdData

var touchCmd = &cobra.Command{
	Use:   "touch FILE",
	Short: "Update the modification time of a file or directory.",
	Long:  `Update the modification time of a file or directory, like touch in Linux.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			err := cmd.Help()
			if err != nil {
				return
			}
			os.Exit(0)
		}

		runTouch(args)
	},
}

func init() {
	rootCmd.AddCommand(touchCmd)
	touchCmd.Flags().BoolVarP(&touchCmdData.isRecursive, "recursive", "r", false, "Update files recursively")
}

func runTouch(args []string) {
	currentTime := time.Now().Local()
	filename := args[0]
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("File not found : %v\n", filename)
	}

	isDir, _ := fileUtil.IsDirectory(filename)
	if isDir && touchCmdData.isRecursive {
		touchRunner.UpdateDirectoryRecursively(filename, currentTime)
	}
	touchRunner.UpdateFile(filename, currentTime)
}
