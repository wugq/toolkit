package cmd

import (
	"fmt"
	"toolkit/runner/iprunner"

	"github.com/spf13/cobra"
)

type IpCmdData struct {
	localOnly  bool
	publicOnly bool
}

var ipCmdData IpCmdData

var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Show IP addresses.",
	Long: `Show local and public IP addresses.

By default shows both local and public IPs.
  -l / --local   Show local IPs only
  -p / --public  Show public IP only

Examples:
  toolkit ip
  toolkit ip -l
  toolkit ip -p`,
	Run: func(cmd *cobra.Command, args []string) {
		runIP()
	},
}

func init() {
	rootCmd.AddCommand(ipCmd)
	ipCmd.Flags().BoolVarP(&ipCmdData.localOnly, "local", "l", false, "Show local IPs only")
	ipCmd.Flags().BoolVarP(&ipCmdData.publicOnly, "public", "p", false, "Show public IP only")
}

func runIP() {
	if ipCmdData.publicOnly {
		ip, err := iprunner.PublicIP()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Public %s\n", ip)
		return
	}
	ips, err := iprunner.LocalIPs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, line := range ips {
		fmt.Println(line)
	}
	if !ipCmdData.localOnly {
		ip, err := iprunner.PublicIP()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Public %s\n", ip)
	}
}
