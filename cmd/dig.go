package cmd

import (
	"fmt"
	"os"
	"toolkit/runner/digRunner"

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
	domain = digRunner.FormatDomain(domain)

	digRunner.PrintAddress(domain)
	digRunner.PrintMX(domain)
	digRunner.PrintTXT(domain)
}
