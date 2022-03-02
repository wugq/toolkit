package digRunner

import (
	"fmt"
	"net"
	"strings"
)

func PrintMX(domain string) {
	mxRecords, _ := net.LookupMX(domain)
	for _, mx := range mxRecords {
		fmt.Printf("%v mail is handled by %v %v\n", domain, mx.Pref, mx.Host)
	}
}

func PrintTXT(domain string) {
	txtRecords, _ := net.LookupTXT(domain)

	for _, txt := range txtRecords {
		fmt.Printf("%v has TXT record %v\n", domain, txt)
	}
}
func PrintAddress(domain string) {
	cname, err := net.LookupCNAME(domain)
	if err != nil {
		fmt.Println(err)
		return
	}

	if cname != domain {
		fmt.Printf("%v is an alias for %v\n", domain, cname)
		PrintAddress(cname)
		return
	}

	ipRecords, _ := net.LookupIP(domain)
	for _, ip := range ipRecords {
		fmt.Printf("%v has address %v\n", domain, ip)
	}
}

func FormatDomain(domain string) string {
	if strings.HasSuffix(domain, ".") {
		return domain
	}
	return domain + "."
}
