package envrunner

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// Env returns all environment variables as formatted lines in a stable, sorted order.
func Env(pretty bool) []string {
	envs := os.Environ()
	sort.Strings(envs)
	var lines []string
	for _, kv := range envs {
		lines = append(lines, formatEnvKV(kv, pretty)...)
	}
	return lines
}

func formatEnvKV(kv string, pretty bool) []string {
	parts := strings.SplitN(kv, "=", 2)
	key := parts[0]
	if len(parts) == 2 {
		value := parts[1]
		if pretty && shouldPrettyPrintValue(key, value) {
			lines := []string{fmt.Sprintf("%s=", key)}
			for _, item := range strings.Split(value, string(os.PathListSeparator)) {
				if item == "" {
					continue
				}
				lines = append(lines, fmt.Sprintf("  - %s", item))
			}
			return lines
		}
	}
	return []string{kv}
}

func shouldPrettyPrintValue(key, value string) bool {
	if value == "" {
		return false
	}
	if strings.ContainsRune(value, os.PathListSeparator) && strings.Contains(key, "PATH") {
		return true
	}
	return false
}
