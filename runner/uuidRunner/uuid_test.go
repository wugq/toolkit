package uuidRunner

import (
	"regexp"
	"testing"
)

var uuidV4Regex = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)

func TestGenerate_Format(t *testing.T) {
	for i := 0; i < 10; i++ {
		got := Generate()
		if !uuidV4Regex.MatchString(got) {
			t.Errorf("Generate() = %q, does not match UUID v4 format", got)
		}
	}
}

func TestGenerate_Unique(t *testing.T) {
	seen := make(map[string]bool)
	for i := 0; i < 100; i++ {
		u := Generate()
		if seen[u] {
			t.Errorf("duplicate UUID generated: %q", u)
		}
		seen[u] = true
	}
}
