package stdin

import (
	"io"
	"os"
)

func IsPiped() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return (stat.Mode() & os.ModeCharDevice) == 0
}

func ReadAll() ([]byte, error) {
	return io.ReadAll(os.Stdin)
}
