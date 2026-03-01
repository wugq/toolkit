package port

import (
	"net"
	"time"
)

// Check returns true if the TCP port is open, false if closed or unreachable.
func Check(host, port string, timeout time.Duration) (bool, error) {
	addr := net.JoinHostPort(host, port)
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return false, nil
	}
	conn.Close()
	return true, nil
}
