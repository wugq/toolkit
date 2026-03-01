package tail

import (
	"bytes"
	"os"
)

// SeekLines returns the file offset at which the last n lines begin.
func SeekLines(fileName string, n int) (int64, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return 0, err
	}

	size := info.Size()
	if size == 0 {
		return 0, nil
	}

	buf := make([]byte, 1)
	newlines := 0
	pos := size - 1

	// Skip a trailing newline (and optional \r) so the last line is counted properly.
	if _, err := file.ReadAt(buf, pos); err == nil && buf[0] == '\n' {
		pos--
		if pos >= 0 {
			if _, err := file.ReadAt(buf, pos); err == nil && buf[0] == '\r' {
				pos--
			}
		}
	}

	for pos >= 0 {
		if _, err := file.ReadAt(buf, pos); err != nil {
			break
		}
		if buf[0] == '\n' {
			newlines++
			if newlines >= n {
				return pos + 1, nil
			}
		}
		pos--
	}
	return 0, nil
}

// ReadFile reads new content between lastPosition and currentPosition.
// Returns the content as a string and the updated position.
func ReadFile(fileName string, lastPosition int64, currentPosition int64) (string, int64) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", lastPosition
	}
	defer file.Close()

	bufSize := currentPosition - lastPosition
	buf := make([]byte, bufSize)
	_, err = file.ReadAt(buf, lastPosition)
	if err == nil && bufSize > 0 {
		// Normalise \r\n to \n so Windows text files display correctly.
		return string(bytes.ReplaceAll(buf, []byte("\r\n"), []byte("\n"))), currentPosition
	}
	return "", currentPosition
}
