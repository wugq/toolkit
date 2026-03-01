package cmd

import (
	"toolkit/runner/ipRunner"

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
		ipRunner.PrintPublicIP()
		return
	}
	ipRunner.PrintLocalIPs()
	if !ipCmdData.localOnly {
		ipRunner.PrintPublicIP()
	}
}
