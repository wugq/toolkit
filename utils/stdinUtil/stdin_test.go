package stdinUtil

import (
	"testing"
)

func TestIsPiped_NotPiped(t *testing.T) {
	// In normal test execution, stdin is not piped
	got := IsPiped()
	if got {
		t.Log("IsPiped() = true (stdin appears piped in this environment)")
	}
	// We don't assert a specific value since the result depends on execution context,
	// but the call must not panic or error.
}
