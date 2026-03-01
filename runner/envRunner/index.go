package envRunner

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// PrintEnv prints all environment variables in a stable, sorted order.
func PrintEnv(pretty bool) {
	envs := os.Environ()
	sort.Strings(envs)
	for _, kv := range envs {
		printEnvKV(kv, pretty)
	}
}

func printEnvKV(kv string, pretty bool) {
	parts := strings.SplitN(kv, "=", 2)
	key := parts[0]
	if len(parts) == 2 {
		value := parts[1]
		if pretty && shouldPrettyPrintValue(key, value) {
			fmt.Printf("%s=\n", key)
			for _, item := range strings.Split(value, string(os.PathListSeparator)) {
				if item == "" {
					continue
				}
				fmt.Printf("  - %s\n", item)
			}
			return
		}
	}
	fmt.Println(kv)
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
