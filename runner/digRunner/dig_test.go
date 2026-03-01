package digRunner

import (
	"net"
	"testing"
)

func TestFormatDomain_AddsTrailingDot(t *testing.T) {
	got := FormatDomain("example.com")
	if got != "example.com." {
		t.Errorf("FormatDomain(%q) = %q, want %q", "example.com", got, "example.com.")
	}
}

func TestFormatDomain_AlreadyHasDot(t *testing.T) {
	got := FormatDomain("example.com.")
	if got != "example.com." {
		t.Errorf("FormatDomain(%q) = %q, want %q", "example.com.", got, "example.com.")
	}
}

func TestFormatDomain_Subdomain(t *testing.T) {
	got := FormatDomain("sub.example.com")
	if got != "sub.example.com." {
		t.Errorf("FormatDomain(%q) = %q, want %q", "sub.example.com", got, "sub.example.com.")
	}
}

func TestNewResolver_EmptyReturnsDefault(t *testing.T) {
	r := NewResolver("")
	if r != net.DefaultResolver {
		t.Error("empty server should return net.DefaultResolver")
	}
}

func TestNewResolver_WithIP(t *testing.T) {
	r := NewResolver("8.8.8.8")
	if r == nil {
		t.Fatal("expected non-nil resolver")
	}
	if r == net.DefaultResolver {
		t.Error("custom server should not return DefaultResolver")
	}
}

func TestNewResolver_WithIPAndPort(t *testing.T) {
	r := NewResolver("8.8.8.8:53")
	if r == nil {
		t.Fatal("expected non-nil resolver")
	}
}
