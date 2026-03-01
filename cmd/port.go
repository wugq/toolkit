package cmd

import (
	"fmt"
	"os"
	"time"
	"toolkit/runner/port"

	"github.com/spf13/cobra"
)

type PortCmdData struct {
	timeout int
}

var portCmdData PortCmdData

var portCmd = &cobra.Command{
	Use:   "port HOST PORT",
	Short: "Check if a TCP port is open.",
	Long: `Check if a TCP port is open on a host.

  --timeout / -T  Connection timeout in seconds (default: 3)

Examples:
  toolkit port example.com 443
  toolkit port 192.168.1.1 22 -T 5`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Please provide a host and port.")
			os.Exit(2)
		}
		timeout := time.Duration(portCmdData.timeout) * time.Second
		open, err := port.Check(args[0], args[1], timeout)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}
		if open {
			fmt.Printf("Port %s on %s is open\n", args[1], args[0])
		} else {
			fmt.Printf("Port %s on %s is closed or unreachable\n", args[1], args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(portCmd)
	portCmd.Flags().IntVarP(&portCmdData.timeout, "timeout", "T", 3, "Connection timeout in seconds")
}
