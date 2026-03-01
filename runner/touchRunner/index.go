package touchrunner

import (
	"os"
	"path/filepath"
	"time"
)

// TouchResult holds the path and type of a touched entry.
type TouchResult struct {
	Path  string
	IsDir bool
}

// UpdateFile updates the modification time of a file or directory.
// Returns the result and any error.
func UpdateFile(fileName string, currentTime time.Time) (TouchResult, error) {
	info, err := os.Stat(fileName)
	if err != nil {
		return TouchResult{}, err
	}
	err = os.Chtimes(fileName, currentTime, currentTime)
	return TouchResult{Path: fileName, IsDir: info.IsDir()}, err
}

// UpdateDirectoryRecursively touches all entries inside dirName recursively.
// Returns all touched results and any error.
func UpdateDirectoryRecursively(dirName string, currentTime time.Time) ([]TouchResult, error) {
	files, err := os.ReadDir(dirName)
	if err != nil {
		return nil, err
	}
	var results []TouchResult
	for _, file := range files {
		filePath := filepath.Join(dirName, file.Name())
		if file.IsDir() {
			subResults, err := UpdateDirectoryRecursively(filePath, currentTime)
			if err != nil {
				return nil, err
			}
			results = append(results, subResults...)
		}
		r, err := UpdateFile(filePath, currentTime)
		if err != nil {
			return nil, err
		}
		results = append(results, r)
	}
	return results, nil
}
