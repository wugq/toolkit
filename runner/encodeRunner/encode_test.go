package encodeRunner

import (
	"testing"
)

func TestBase64Encode_Decode_RoundTrip(t *testing.T) {
	inputs := []string{
		"",
		"hello",
		"hello world",
		"special: !@#$%^&*()",
		"multi\nline\ntext",
	}
	for _, input := range inputs {
		t.Run(input, func(t *testing.T) {
			encoded := Base64Encode(input)
			decoded, err := Base64Decode(encoded)
			if err != nil {
				t.Fatalf("Base64Decode error: %v", err)
			}
			if decoded != input {
				t.Errorf("round-trip failed: got %q, want %q", decoded, input)
			}
		})
	}
}

func TestBase64Encode_KnownValue(t *testing.T) {
	got := Base64Encode("hello")
	want := "aGVsbG8="
	if got != want {
		t.Errorf("Base64Encode(%q) = %q, want %q", "hello", got, want)
	}
}

func TestBase64Decode_Invalid(t *testing.T) {
	_, err := Base64Decode("not-valid-base64!!!")
	if err == nil {
		t.Error("expected error for invalid base64 input")
	}
}

func TestURLEncode_Decode_RoundTrip(t *testing.T) {
	inputs := []string{
		"hello world",
		"foo=bar&baz=qux",
		"https://example.com/path?q=hello world",
		"special: !@#$%",
	}
	for _, input := range inputs {
		t.Run(input, func(t *testing.T) {
			encoded := URLEncode(input)
			decoded, err := URLDecode(encoded)
			if err != nil {
				t.Fatalf("URLDecode error: %v", err)
			}
			if decoded != input {
				t.Errorf("round-trip failed: got %q, want %q", decoded, input)
			}
		})
	}
}

func TestURLEncode_EncodesSpaces(t *testing.T) {
	got := URLEncode("hello world")
	if got != "hello+world" {
		t.Errorf("URLEncode(%q) = %q, want %q", "hello world", got, "hello+world")
	}
}

func TestURLDecode_Invalid(t *testing.T) {
	_, err := URLDecode("hello%ZZ")
	if err == nil {
		t.Error("expected error for invalid URL-encoded input")
	}
}
