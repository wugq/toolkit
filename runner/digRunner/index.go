package digRunner

import (
	"context"
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

func NS(domain string, r *net.Resolver) ([]*net.NS, error) {
	return r.LookupNS(context.Background(), domain)
}

func CNAME(domain string, r *net.Resolver) (string, error) {
	return r.LookupCNAME(context.Background(), domain)
}

func MX(domain string, r *net.Resolver) ([]*net.MX, error) {
	return r.LookupMX(context.Background(), domain)
}

func TXT(domain string, r *net.Resolver) ([]string, error) {
	return r.LookupTXT(context.Background(), domain)
}

// AliasEntry holds one step in a CNAME alias chain.
type AliasEntry struct {
	From string
	To   string
}

// AddressResult holds the resolved address information for a domain.
type AddressResult struct {
	Aliases []AliasEntry
	Final   string
	IPs     []net.IP
}

// Address resolves a domain, following CNAME aliases, and returns all IPs.
func Address(domain string, r *net.Resolver) (AddressResult, error) {
	var result AddressResult
	current := domain
	for {
		cname, err := r.LookupCNAME(context.Background(), current)
		if err != nil {
			return AddressResult{}, err
		}
		if cname == current {
			result.Final = current
			break
		}
		result.Aliases = append(result.Aliases, AliasEntry{From: current, To: cname})
		current = cname
		if len(result.Aliases) > 20 {
			result.Final = current
			break
		}
	}
	result.IPs, _ = r.LookupIP(context.Background(), "ip", result.Final)
	return result, nil
}

func FormatDomain(domain string) string {
	if strings.HasSuffix(domain, ".") {
		return domain
	}
	return domain + "."
}
