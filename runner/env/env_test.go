package env

import (
	"os"
	"strings"
	"testing"
)

func TestFormatEnvKV_Simple(t *testing.T) {
	lines := formatEnvKV("FOO=bar", false)
	if len(lines) != 1 || lines[0] != "FOO=bar" {
		t.Errorf("unexpected result: %v", lines)
	}
}

func TestFormatEnvKV_NoValue(t *testing.T) {
	lines := formatEnvKV("NOVALUE", false)
	if len(lines) != 1 || lines[0] != "NOVALUE" {
		t.Errorf("unexpected result for key-only entry: %v", lines)
	}
}

func TestFormatEnvKV_Pretty_PATH(t *testing.T) {
	sep := string(os.PathListSeparator)
	kv := "PATH=/usr/bin" + sep + "/usr/local/bin"
	lines := formatEnvKV(kv, true)
	if len(lines) < 2 {
		t.Fatalf("expected multiple lines for pretty PATH, got: %v", lines)
	}
	if lines[0] != "PATH=" {
		t.Errorf("expected header line %q, got %q", "PATH=", lines[0])
	}
	for _, line := range lines[1:] {
		if !strings.HasPrefix(line, "  - ") {
			t.Errorf("expected indented item, got %q", line)
		}
	}
}

func TestFormatEnvKV_NoPretty_PATH(t *testing.T) {
	sep := string(os.PathListSeparator)
	kv := "PATH=/usr/bin" + sep + "/usr/local/bin"
	lines := formatEnvKV(kv, false)
	if len(lines) != 1 || lines[0] != kv {
		t.Errorf("expected single unchanged line, got: %v", lines)
	}
}

func TestFormatEnvKV_Pretty_NonPATH(t *testing.T) {
	sep := string(os.PathListSeparator)
	kv := "MY_VAR=a" + sep + "b"
	lines := formatEnvKV(kv, true)
	// should not pretty-print since key doesn't contain "PATH"
	if len(lines) != 1 || lines[0] != kv {
		t.Errorf("expected single line for non-PATH var, got: %v", lines)
	}
}

func TestEnv_IsSorted(t *testing.T) {
	os.Setenv("ZZZ_TEST_TOOLKIT", "1")
	os.Setenv("AAA_TEST_TOOLKIT", "2")
	defer os.Unsetenv("ZZZ_TEST_TOOLKIT")
	defer os.Unsetenv("AAA_TEST_TOOLKIT")

	lines := Env(false)
	for i := 1; i < len(lines); i++ {
		if lines[i] < lines[i-1] {
			t.Errorf("output not sorted: %q before %q", lines[i-1], lines[i])
		}
	}
}

func TestEnv_ContainsSetVars(t *testing.T) {
	os.Setenv("ZZZ_TEST_TOOLKIT", "hello")
	defer os.Unsetenv("ZZZ_TEST_TOOLKIT")

	lines := Env(false)
	found := false
	for _, line := range lines {
		if line == "ZZZ_TEST_TOOLKIT=hello" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected ZZZ_TEST_TOOLKIT=hello in output")
	}
}
