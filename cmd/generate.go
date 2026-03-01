package cmd

import (
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate TYPE",
	Short: "Generate things like password, UUID and so on",
	Long: `Generate random values.

Subcommands:
  password  Generate a random password with configurable character sets and length
  uuid      Generate a random UUID (v4)

Run 'toolkit generate <subcommand> --help' for details.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
