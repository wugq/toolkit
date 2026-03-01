package json

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(input []byte) (string, error) {
	var v interface{}
	if err := json.Unmarshal(input, &v); err != nil {
		return "", fmt.Errorf("invalid JSON: %v", err)
	}
	out, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", err
	}
	return string(out), nil
}
