package cmd

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var digCmd = &cobra.Command{
	Use:   "dig domain",
	Short: "DNS lookup.",
	Long:  `DNS lookup.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Please enter a domain.")
			os.Exit(2)
		}
		runDig(args)
	},
}

func init() {
	rootCmd.AddCommand(digCmd)
}

func runDig(args []string) {
	var domain = args[0]
	domain = formatDomain(domain)

	printAddress(domain)
	printMX(domain)
	printTXT(domain)
}

func printMX(domain string) {
	mxRecords, _ := net.LookupMX(domain)
	for _, mx := range mxRecords {
		fmt.Printf("%v mail is handled by %v %v\n", domain, mx.Pref, mx.Host)
	}
}

func printTXT(domain string) {
	txtRecords, _ := net.LookupTXT(domain)

	for _, txt := range txtRecords {
		fmt.Printf("%v has TXT record %v\n", domain, txt)
	}
}
func printAddress(domain string) {
	cname, err := net.LookupCNAME(domain)
	if err != nil {
		fmt.Println(err)
		return
	}

	if cname != domain {
		fmt.Printf("%v is an alias for %v\n", domain, cname)
		printAddress(cname)
		return
	}

	ipRecords, _ := net.LookupIP(domain)
	for _, ip := range ipRecords {
		fmt.Printf("%v has address %v\n", domain, ip)
	}
}

func formatDomain(domain string) string {
	if strings.HasSuffix(domain, ".") {
		return domain
	}
	return domain + "."
}
