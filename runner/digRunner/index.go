package digRunner

import (
	"context"
	"fmt"
	"net"
	"strings"
)

func NewResolver(server string) *net.Resolver {
	if server == "" {
		return net.DefaultResolver
	}
	addr := server
	if !strings.Contains(addr, ":") {
		addr = addr + ":53"
	}
	return &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{}
			return d.DialContext(ctx, "udp", addr)
		},
	}
}

func PrintNS(domain string, r *net.Resolver) {
	nsRecords, err := r.LookupNS(context.Background(), domain)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, ns := range nsRecords {
		fmt.Printf("%v name server %v\n", domain, ns.Host)
	}
}

func PrintCNAME(domain string, r *net.Resolver) {
	cname, err := r.LookupCNAME(context.Background(), domain)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v canonical name %v\n", domain, cname)
}

func PrintMX(domain string, r *net.Resolver) {
	mxRecords, _ := r.LookupMX(context.Background(), domain)
	for _, mx := range mxRecords {
		fmt.Printf("%v mail is handled by %v %v\n", domain, mx.Pref, mx.Host)
	}
}

func PrintTXT(domain string, r *net.Resolver) {
	txtRecords, _ := r.LookupTXT(context.Background(), domain)
	for _, txt := range txtRecords {
		fmt.Printf("%v has TXT record %v\n", domain, txt)
	}
}

func PrintAddress(domain string, r *net.Resolver) {
	cname, err := r.LookupCNAME(context.Background(), domain)
	if err != nil {
		fmt.Println(err)
		return
	}

	if cname != domain {
		fmt.Printf("%v is an alias for %v\n", domain, cname)
		PrintAddress(cname, r)
		return
	}

	ipRecords, _ := r.LookupIP(context.Background(), "ip", domain)
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
