package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "toolkit",
	Short: "WuGQ's toolkit",
	Long: `WuGQ's toolkit — a collection of everyday CLI utilities.

Commands:
  date    Show or convert dates and timestamps, with date arithmetic
  dig     DNS lookup for A, CNAME, MX, NS, and TXT records
  md5sum  Compute the MD5 checksum of a file or text string
  tail    Print the end of a file, with optional follow mode
  touch   Update the modification time of a file or directory
  env     List environment variables

  generate password  Generate a random password
  generate uuid      Generate a random UUID

Run 'toolkit <command> --help' for details on each command.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
