package jsonrunner

import (
	"strings"
	"testing"
)

func TestPrettyPrint_Object(t *testing.T) {
	input := `{"b":2,"a":1}`
	out, err := PrettyPrint([]byte(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "\n") {
		t.Error("expected formatted JSON with newlines")
	}
	if !strings.Contains(out, `"a"`) || !strings.Contains(out, `"b"`) {
		t.Error("output should contain original keys")
	}
}

func TestPrettyPrint_Array(t *testing.T) {
	input := `[1,2,3]`
	out, err := PrettyPrint([]byte(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "1") || !strings.Contains(out, "2") || !strings.Contains(out, "3") {
		t.Error("output should contain array elements")
	}
}

func TestPrettyPrint_Nested(t *testing.T) {
	input := `{"x":{"y":{"z":42}}}`
	out, err := PrettyPrint([]byte(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "42") {
		t.Error("output should contain nested value")
	}
}

func TestPrettyPrint_Invalid(t *testing.T) {
	_, err := PrettyPrint([]byte("not json"))
	if err == nil {
		t.Error("expected error for invalid JSON")
	}
}

func TestPrettyPrint_InvalidPartial(t *testing.T) {
	_, err := PrettyPrint([]byte(`{"key": `))
	if err == nil {
		t.Error("expected error for partial JSON")
	}
}

func TestPrettyPrint_Empty(t *testing.T) {
	_, err := PrettyPrint([]byte(""))
	if err == nil {
		t.Error("expected error for empty input")
	}
}

func TestPrettyPrint_String(t *testing.T) {
	out, err := PrettyPrint([]byte(`"hello"`))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out != `"hello"` {
		t.Errorf("got %q, want %q", out, `"hello"`)
	}
}
