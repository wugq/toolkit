package tail

import (
	"os"
	"testing"
)

func writeTempFile(t *testing.T, content string) string {
	t.Helper()
	f, err := os.CreateTemp("", "tailtest")
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.WriteString(content)
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	return f.Name()
}

func TestSeekLines_Last1(t *testing.T) {
	path := writeTempFile(t, "line1\nline2\nline3\n")
	defer os.Remove(path)

	offset, err := SeekLines(path, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	content := "line1\nline2\nline3\n"
	got := content[offset:]
	if got != "line3\n" {
		t.Errorf("got %q, want %q", got, "line3\n")
	}
}

func TestSeekLines_Last2(t *testing.T) {
	path := writeTempFile(t, "line1\nline2\nline3\n")
	defer os.Remove(path)

	offset, err := SeekLines(path, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	content := "line1\nline2\nline3\n"
	got := content[offset:]
	if got != "line2\nline3\n" {
		t.Errorf("got %q, want %q", got, "line2\nline3\n")
	}
}

func TestSeekLines_MoreThanAvailable(t *testing.T) {
	content := "line1\nline2\n"
	path := writeTempFile(t, content)
	defer os.Remove(path)

	offset, err := SeekLines(path, 100)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if offset != 0 {
		t.Errorf("expected offset 0 when requesting more lines than available, got %d", offset)
	}
}

func TestSeekLines_EmptyFile(t *testing.T) {
	path := writeTempFile(t, "")
	defer os.Remove(path)

	offset, err := SeekLines(path, 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if offset != 0 {
		t.Errorf("expected offset 0 for empty file, got %d", offset)
	}
}

func TestSeekLines_NotFound(t *testing.T) {
	_, err := SeekLines("/nonexistent/file.txt", 5)
	if err == nil {
		t.Error("expected error for nonexistent file")
	}
}

func TestReadFile_FullContent(t *testing.T) {
	content := "hello\nworld\n"
	path := writeTempFile(t, content)
	defer os.Remove(path)

	got, newPos := ReadFile(path, 0, int64(len(content)))
	if got != content {
		t.Errorf("got %q, want %q", got, content)
	}
	if newPos != int64(len(content)) {
		t.Errorf("newPos = %d, want %d", newPos, len(content))
	}
}

func TestReadFile_Partial(t *testing.T) {
	content := "hello\nworld\n"
	path := writeTempFile(t, content)
	defer os.Remove(path)

	got, _ := ReadFile(path, 6, int64(len(content)))
	if got != "world\n" {
		t.Errorf("got %q, want %q", got, "world\n")
	}
}

func TestReadFile_CRLF_Normalisation(t *testing.T) {
	content := "line1\r\nline2\r\n"
	path := writeTempFile(t, content)
	defer os.Remove(path)

	got, _ := ReadFile(path, 0, int64(len(content)))
	want := "line1\nline2\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestReadFile_ZeroRange(t *testing.T) {
	content := "hello"
	path := writeTempFile(t, content)
	defer os.Remove(path)

	got, newPos := ReadFile(path, 3, 3)
	if got != "" {
		t.Errorf("expected empty string for zero range, got %q", got)
	}
	if newPos != 3 {
		t.Errorf("newPos = %d, want 3", newPos)
	}
}

func TestReadFile_NotFound(t *testing.T) {
	got, pos := ReadFile("/nonexistent/file.txt", 0, 10)
	if got != "" {
		t.Errorf("expected empty string for missing file, got %q", got)
	}
	if pos != 0 {
		t.Errorf("expected lastPosition returned on error, got %d", pos)
	}
}
