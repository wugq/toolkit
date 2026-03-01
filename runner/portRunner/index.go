package portRunner

import (
	"fmt"
	"net"
	"time"
)

func Check(host, port string, timeout time.Duration) {
	addr := net.JoinHostPort(host, port)
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		fmt.Printf("Port %s on %s is closed or unreachable\n", port, host)
		return
	}
	conn.Close()
	fmt.Printf("Port %s on %s is open\n", port, host)
}
