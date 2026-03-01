package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"
	"toolkit/runner/touch"
	"toolkit/util/file"
)

type TouchCmdData struct {
	isRecursive bool
}

var touchCmdData TouchCmdData

var touchCmd = &cobra.Command{
	Use:   "touch FILE",
	Short: "Update the modification time of a file or directory.",
	Long: `Update the modification time of a file or directory to the current time.

  -r / --recursive  Also update all files inside a directory recursively

Examples:
  toolkit touch ./file.txt
  toolkit touch -r ./some-dir`,
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

	isDir, _ := file.IsDirectory(filename)
	if isDir && touchCmdData.isRecursive {
		results, err := touch.UpdateDirectoryRecursively(filename, currentTime)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, r := range results {
			if r.IsDir {
				fmt.Printf("Touch directory: %v\n", r.Path)
			} else {
				fmt.Printf("Touch file: %v\n", r.Path)
			}
		}
	}
	r, err := touch.UpdateFile(filename, currentTime)
	if err != nil {
		fmt.Println(err)
		return
	}
	if r.IsDir {
		fmt.Printf("Touch directory: %v\n", r.Path)
	} else {
		fmt.Printf("Touch file: %v\n", r.Path)
	}
}
