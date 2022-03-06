package md5sumRunner

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func CheckFile(filePath string) (string, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return "", err
	}

	defer file.Close()

	hash := md5.New()
	_, err = io.Copy(hash, file)

	if err != nil {
		return "", err
	}
	md5sum := hash.Sum(nil)
	return fmt.Sprintf("%x", md5sum), nil
}

func CheckText(message string) (string, error) {
	res := md5.Sum([]byte(message))
	return fmt.Sprintf("%x", res), nil

}
