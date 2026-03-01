package portRunner

import (
	"net"
	"testing"
	"time"
)

func TestCheck_OpenPort(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	host, port, _ := net.SplitHostPort(ln.Addr().String())
	open, err := Check(host, port, 2*time.Second)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !open {
		t.Error("expected port to be open")
	}
}

func TestCheck_ClosedPort(t *testing.T) {
	// Port 1 requires root to listen, so it should be closed
	open, err := Check("127.0.0.1", "1", 500*time.Millisecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if open {
		t.Error("expected port 1 to be closed")
	}
}

func TestCheck_InvalidHost(t *testing.T) {
	open, err := Check("invalid.host.that.does.not.exist.local", "80", 500*time.Millisecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if open {
		t.Error("expected unreachable host to return closed")
	}
}
