package token

import (
	"encoding/hex"
	"testing"
)

func TestGenerate_Length(t *testing.T) {
	for _, length := range []int{8, 16, 32, 64} {
		t.Run("length", func(t *testing.T) {
			token, err := Generate(length)
			if err != nil {
				t.Fatalf("Generate(%d) error: %v", length, err)
			}
			// hex encoding doubles the byte count
			if len(token) != length*2 {
				t.Errorf("Generate(%d) output length = %d, want %d", length, len(token), length*2)
			}
		})
	}
}

func TestGenerate_ValidHex(t *testing.T) {
	token, err := Generate(16)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := hex.DecodeString(token); err != nil {
		t.Errorf("Generate output %q is not valid hex: %v", token, err)
	}
}

func TestGenerate_Unique(t *testing.T) {
	t1, err := Generate(16)
	if err != nil {
		t.Fatal(err)
	}
	t2, err := Generate(16)
	if err != nil {
		t.Fatal(err)
	}
	if t1 == t2 {
		t.Error("two generated tokens should not be equal")
	}
}

func TestGenerate_Zero(t *testing.T) {
	token, err := Generate(0)
	if err != nil {
		t.Fatal(err)
	}
	if token != "" {
		t.Errorf("Generate(0) = %q, want empty string", token)
	}
}
