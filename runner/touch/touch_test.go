package touch

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestUpdateFile_File(t *testing.T) {
	f, err := os.CreateTemp("", "touchtest")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	defer os.Remove(f.Name())

	now := time.Now().Truncate(time.Second)
	result, err := UpdateFile(f.Name(), now)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.IsDir {
		t.Error("expected IsDir=false for a file")
	}
	if result.Path != f.Name() {
		t.Errorf("Path = %q, want %q", result.Path, f.Name())
	}

	info, err := os.Stat(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	if !info.ModTime().Truncate(time.Second).Equal(now) {
		t.Errorf("mod time = %v, want %v", info.ModTime().Truncate(time.Second), now)
	}
}

func TestUpdateFile_Directory(t *testing.T) {
	dir, err := os.MkdirTemp("", "touchdirtest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	now := time.Now().Truncate(time.Second)
	result, err := UpdateFile(dir, now)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsDir {
		t.Error("expected IsDir=true for a directory")
	}
	if result.Path != dir {
		t.Errorf("Path = %q, want %q", result.Path, dir)
	}
}

func TestUpdateFile_NotFound(t *testing.T) {
	_, err := UpdateFile("/nonexistent/path/file.txt", time.Now())
	if err == nil {
		t.Error("expected error for nonexistent path")
	}
}

func TestUpdateDirectoryRecursively(t *testing.T) {
	dir, err := os.MkdirTemp("", "touchrectest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	// Create subdir with a file inside
	subdir := filepath.Join(dir, "sub")
	if err := os.Mkdir(subdir, 0755); err != nil {
		t.Fatal(err)
	}
	f, err := os.Create(filepath.Join(subdir, "file.txt"))
	if err != nil {
		t.Fatal(err)
	}
	f.Close()

	now := time.Now().Truncate(time.Second)
	results, err := UpdateDirectoryRecursively(dir, now)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Expect: sub/file.txt and sub/ itself
	if len(results) < 2 {
		t.Errorf("expected at least 2 results, got %d", len(results))
	}
}

func TestUpdateDirectoryRecursively_EmptyDir(t *testing.T) {
	dir, err := os.MkdirTemp("", "touchrectest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	results, err := UpdateDirectoryRecursively(dir, time.Now())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 0 {
		t.Errorf("expected 0 results for empty dir, got %d", len(results))
	}
}

func TestUpdateDirectoryRecursively_NotFound(t *testing.T) {
	_, err := UpdateDirectoryRecursively("/nonexistent/path", time.Now())
	if err == nil {
		t.Error("expected error for nonexistent directory")
	}
}
