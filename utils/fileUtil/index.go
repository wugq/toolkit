package fileUtil

import "os"

func IsDirectory(directoryName string) (bool, error) {
	fileInfo, err := os.Stat(directoryName)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), err
}
