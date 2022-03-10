package touchRunner

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

func UpdateFile(fileName string, currentTime time.Time) {
	file, err := os.Stat(fileName)

	if err != nil {
		fmt.Println(err)
		return
	}

	if file.IsDir() {
		fmt.Printf("Touch directory: %v\n", fileName)
	} else {
		fmt.Printf("Touch file: %v\n", fileName)
	}

	err = os.Chtimes(fileName, currentTime, currentTime)

}

func UpdateDirectoryRecursively(dirName string, currentTime time.Time) {
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		var filePath = filepath.Join(dirName, file.Name())
		if file.IsDir() {
			UpdateDirectoryRecursively(filePath, currentTime)
		}
		UpdateFile(filePath, currentTime)
	}
}
