package fileUtil

import "os"

func IsDirectory(directoryName string) (bool, error) {
	fileInfo, err := os.Stat(directoryName)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), err
}

func GetFileSize(fileName string) (int64, error) {
	stat, err := os.Stat(fileName)
	return stat.Size(), err
}

func IsFile(fileName string) (bool, error) {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		return false, err
	}

	return !fileInfo.IsDir(), err
}
