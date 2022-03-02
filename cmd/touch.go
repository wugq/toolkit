package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

type TouchCmdData struct {
	isRecursive bool
	currentTime time.Time
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
	touchCmdData.currentTime = time.Now().Local()
	var filename = args[0]
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("File not found : %v\n", filename)
	}

	var isDir, _ = isDirectory(filename)
	if isDir && touchCmdData.isRecursive {
		updateDirectory(filename)
	}
	touchFile(filename)
}

func touchFile(filename string) {
	file, err := os.Stat(filename)

	if err != nil {
		fmt.Println(err)
		return
	}

	if file.IsDir() {
		fmt.Printf("Touch directory: %v\n", filename)
	} else {
		fmt.Printf("Touch file: %v\n", filename)
	}

	err = os.Chtimes(filename, touchCmdData.currentTime, touchCmdData.currentTime)

}

func updateDirectory(directoryName string) {
	files, err := ioutil.ReadDir(directoryName)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		var filePath = filepath.Join(directoryName, file.Name())
		if file.IsDir() {
			updateDirectory(filePath)
		}
		touchFile(filePath)
	}
}

func isDirectory(directoryName string) (bool, error) {
	fileInfo, err := os.Stat(directoryName)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), err
}
