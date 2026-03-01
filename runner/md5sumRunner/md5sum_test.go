package md5sumrunner

import (
	"encoding/hex"
	"os"
	"testing"
)

func TestCheckText_MD5_KnownValues(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"", "d41d8cd98f00b204e9800998ecf8427e"},
		{"hello", "5d41402abc4b2a76b9719d911017c592"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := CheckText(tt.input, "md5")
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("MD5(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestCheckText_OutputLength(t *testing.T) {
	tests := []struct {
		algo    string
		wantLen int
	}{
		{"md5", 32},
		{"sha256", 64},
		{"sha512", 128},
	}
	for _, tt := range tests {
		t.Run(tt.algo, func(t *testing.T) {
			got, err := CheckText("hello", tt.algo)
			if err != nil {
				t.Fatalf("CheckText(\"hello\", %q) error: %v", tt.algo, err)
			}
			if len(got) != tt.wantLen {
				t.Errorf("CheckText(\"hello\", %q) length = %d, want %d", tt.algo, len(got), tt.wantLen)
			}
			if _, err := hex.DecodeString(got); err != nil {
				t.Errorf("CheckText output is not valid hex: %v", err)
			}
		})
	}
}

func TestCheckText_Consistency(t *testing.T) {
	for _, algo := range []string{"md5", "sha256", "sha512"} {
		a, err := CheckText("consistent input", algo)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		b, _ := CheckText("consistent input", algo)
		if a != b {
			t.Errorf("%s: same input produced different hashes", algo)
		}
	}
}

func TestCheckText_DifferentInputs(t *testing.T) {
	for _, algo := range []string{"md5", "sha256", "sha512"} {
		a, _ := CheckText("foo", algo)
		b, _ := CheckText("bar", algo)
		if a == b {
			t.Errorf("%s: different inputs produced same hash", algo)
		}
	}
}

func TestCheckText_UnsupportedAlgo(t *testing.T) {
	_, err := CheckText("hello", "crc32")
	if err == nil {
		t.Error("expected error for unsupported algorithm")
	}
}

func TestCheckFile_MD5(t *testing.T) {
	f, err := os.CreateTemp("", "md5sumtest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())
	f.WriteString("hello")
	f.Close()

	got, err := CheckFile(f.Name(), "md5")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "5d41402abc4b2a76b9719d911017c592"
	if got != want {
		t.Errorf("CheckFile md5 = %q, want %q", got, want)
	}
}

func TestCheckFile_MatchesCheckText(t *testing.T) {
	content := "test file content for hashing"
	f, err := os.CreateTemp("", "md5sumtest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())
	f.WriteString(content)
	f.Close()

	for _, algo := range []string{"md5", "sha256", "sha512"} {
		fileHash, err := CheckFile(f.Name(), algo)
		if err != nil {
			t.Fatalf("CheckFile error for %s: %v", algo, err)
		}
		textHash, err := CheckText(content, algo)
		if err != nil {
			t.Fatalf("CheckText error for %s: %v", algo, err)
		}
		if fileHash != textHash {
			t.Errorf("%s: CheckFile = %q, CheckText = %q (should match)", algo, fileHash, textHash)
		}
	}
}

func TestCheckFile_NotFound(t *testing.T) {
	_, err := CheckFile("/nonexistent/path/file.txt", "md5")
	if err == nil {
		t.Error("expected error for nonexistent file")
	}
}
