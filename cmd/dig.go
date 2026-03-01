package cmd

import (
	"fmt"
	"os"
	"strings"
	"toolkit/runner/digRunner"

	"github.com/spf13/cobra"
)

type DigCmdData struct {
	RecordType string
	Verbose    bool
	Resolver   string
}

var digCmdData DigCmdData

var digCmd = &cobra.Command{
	Use:   "dig domain",
	Short: "DNS lookup.",
	Long: `Look up DNS records for a domain.

By default shows A and CNAME records. Use flags to see more:
  -v / --verbose  Also show MX, NS, and TXT records
  -t / --type     Query a specific record type: A, CNAME, MX, NS, TXT

Examples:
  toolkit dig example.com
  toolkit dig example.com -v
  toolkit dig example.com -t MX`,
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
	digCmd.Flags().StringVarP(&digCmdData.RecordType, "type", "t", "", "record type to query: A, CNAME, MX, NS, TXT")
	digCmd.Flags().BoolVarP(&digCmdData.Verbose, "verbose", "v", false, "also show MX, NS, and TXT records")
	digCmd.Flags().StringVarP(&digCmdData.Resolver, "resolver", "r", "", "DNS server to use (e.g. 8.8.8.8)")
}

func runDig(args []string) {
	domain := digRunner.FormatDomain(args[0])
	recordType := strings.ToUpper(digCmdData.RecordType)
	resolver := digRunner.NewResolver(digCmdData.Resolver)

	if recordType != "" {
		switch recordType {
		case "A":
			digRunner.PrintAddress(domain, resolver)
		case "CNAME":
			digRunner.PrintCNAME(domain, resolver)
		case "MX":
			digRunner.PrintMX(domain, resolver)
		case "NS":
			digRunner.PrintNS(domain, resolver)
		case "TXT":
			digRunner.PrintTXT(domain, resolver)
		default:
			fmt.Printf("Unsupported record type: %s\nSupported types: A, CNAME, MX, NS, TXT\n", digCmdData.RecordType)
			os.Exit(2)
		}
		return
	}

	digRunner.PrintAddress(domain, resolver)
	if digCmdData.Verbose {
		digRunner.PrintMX(domain, resolver)
		digRunner.PrintNS(domain, resolver)
		digRunner.PrintTXT(domain, resolver)
	}
}
