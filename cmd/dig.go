package cmd

import (
	"fmt"
	"os"
	"strings"
	"toolkit/runner/dig"

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
	domain := dig.FormatDomain(args[0])
	recordType := strings.ToUpper(digCmdData.RecordType)
	resolver := dig.NewResolver(digCmdData.Resolver)

	if recordType != "" {
		switch recordType {
		case "A":
			result, err := dig.Address(domain, resolver)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			printAddress(result)
		case "CNAME":
			cname, err := dig.CNAME(domain, resolver)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Printf("%v canonical name %v\n", domain, cname)
		case "MX":
			records, err := dig.MX(domain, resolver)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			for _, mx := range records {
				fmt.Printf("%v mail is handled by %v %v\n", domain, mx.Pref, mx.Host)
			}
		case "NS":
			records, err := dig.NS(domain, resolver)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			for _, ns := range records {
				fmt.Printf("%v name server %v\n", domain, ns.Host)
			}
		case "TXT":
			records, err := dig.TXT(domain, resolver)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			for _, txt := range records {
				fmt.Printf("%v has TXT record %v\n", domain, txt)
			}
		default:
			fmt.Printf("Unsupported record type: %s\nSupported types: A, CNAME, MX, NS, TXT\n", digCmdData.RecordType)
			os.Exit(2)
		}
		return
	}

	result, err := dig.Address(domain, resolver)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	printAddress(result)
	if digCmdData.Verbose {
		if records, err := dig.MX(domain, resolver); err == nil {
			for _, mx := range records {
				fmt.Printf("%v mail is handled by %v %v\n", domain, mx.Pref, mx.Host)
			}
		}
		if records, err := dig.NS(domain, resolver); err == nil {
			for _, ns := range records {
				fmt.Printf("%v name server %v\n", domain, ns.Host)
			}
		}
		if records, err := dig.TXT(domain, resolver); err == nil {
			for _, txt := range records {
				fmt.Printf("%v has TXT record %v\n", domain, txt)
			}
		}
	}
}

func printAddress(result dig.AddressResult) {
	for _, alias := range result.Aliases {
		fmt.Printf("%v is an alias for %v\n", alias.From, alias.To)
	}
	for _, ip := range result.IPs {
		fmt.Printf("%v has address %v\n", result.Final, ip)
	}
}
