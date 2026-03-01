package md5sum

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"io"
	"os"
	"strings"
)

func newHash(algo string) (hash.Hash, error) {
	switch strings.ToLower(algo) {
	case "md5":
		return md5.New(), nil
	case "sha256":
		return sha256.New(), nil
	case "sha512":
		return sha512.New(), nil
	default:
		return nil, fmt.Errorf("unsupported algorithm %q: choose md5, sha256, or sha512", algo)
	}
}

func CheckFile(filePath string, algo string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	h, err := newHash(algo)
	if err != nil {
		return "", err
	}

	if _, err = io.Copy(h, file); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func CheckText(message string, algo string) (string, error) {
	h, err := newHash(algo)
	if err != nil {
		return "", err
	}
	h.Write([]byte(message))
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
