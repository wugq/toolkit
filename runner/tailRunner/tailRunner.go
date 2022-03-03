package tailRunner

import (
	"fmt"
	"os"
)

func ReadFile(fileName string, lastPosition int64, currentPosition int64) int64 {
	file, err := os.Open(fileName)
	if err != nil {
		return lastPosition
	}
	defer file.Close()

	bufSize := currentPosition - lastPosition
	buf := make([]byte, bufSize)
	_, err = file.ReadAt(buf, lastPosition)
	if err == nil && bufSize > 0 {
		fmt.Printf("%s", buf)
	}
	return currentPosition

}
