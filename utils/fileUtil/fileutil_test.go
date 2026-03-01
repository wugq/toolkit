package fileUtil

import (
	"os"
	"testing"
)

func TestIsDirectory_WithDir(t *testing.T) {
	dir, err := os.MkdirTemp("", "fileutiltest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	got, err := IsDirectory(dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !got {
		t.Errorf("IsDirectory(%q) = false, want true", dir)
	}
}

func TestIsDirectory_WithFile(t *testing.T) {
	f, err := os.CreateTemp("", "fileutiltest")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	defer os.Remove(f.Name())

	got, err := IsDirectory(f.Name())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got {
		t.Errorf("IsDirectory(%q) = true for a file, want false", f.Name())
	}
}

func TestIsDirectory_NotFound(t *testing.T) {
	_, err := IsDirectory("/nonexistent/path")
	if err == nil {
		t.Error("expected error for nonexistent path")
	}
}

func TestIsFile_WithFile(t *testing.T) {
	f, err := os.CreateTemp("", "fileutiltest")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	defer os.Remove(f.Name())

	got, err := IsFile(f.Name())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !got {
		t.Errorf("IsFile(%q) = false, want true", f.Name())
	}
}

func TestIsFile_WithDir(t *testing.T) {
	dir, err := os.MkdirTemp("", "fileutiltest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	got, err := IsFile(dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got {
		t.Errorf("IsFile(%q) = true for a directory, want false", dir)
	}
}

func TestIsFile_NotFound(t *testing.T) {
	_, err := IsFile("/nonexistent/file.txt")
	if err == nil {
		t.Error("expected error for nonexistent path")
	}
}

func TestGetFileSize(t *testing.T) {
	f, err := os.CreateTemp("", "fileutiltest")
	if err != nil {
		t.Fatal(err)
	}
	_, _ = f.WriteString("hello")
	f.Close()
	defer os.Remove(f.Name())

	size, err := GetFileSize(f.Name())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if size != 5 {
		t.Errorf("GetFileSize = %d, want 5", size)
	}
}

func TestGetFileSize_Empty(t *testing.T) {
	f, err := os.CreateTemp("", "fileutiltest")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	defer os.Remove(f.Name())

	size, err := GetFileSize(f.Name())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if size != 0 {
		t.Errorf("GetFileSize = %d, want 0", size)
	}
}
