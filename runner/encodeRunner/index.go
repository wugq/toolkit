package encoderunner

import (
	"encoding/base64"
	"net/url"
)

func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func Base64Decode(s string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func URLEncode(s string) string {
	return url.QueryEscape(s)
}

func URLDecode(s string) (string, error) {
	return url.QueryUnescape(s)
}
