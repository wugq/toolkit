package ipRunner

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

func PrintLocalIPs() {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			fmt.Printf("Local  [%s] %s\n", iface.Name, ip.String())
		}
	}
}

func PrintPublicIP() {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		fmt.Printf("Could not fetch public IP: %v\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Could not read response: %v\n", err)
		return
	}
	fmt.Printf("Public %s\n", strings.TrimSpace(string(body)))
}
