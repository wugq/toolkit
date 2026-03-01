package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// Example for Flags
// var Verbose bool
// var Source string

// rootCmd represents the base command when called without any subcommands
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

  generate password  Generate a random password
  generate uuid      Generate a random UUID

Run 'toolkit <command> --help' for details on each command.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config fileUtil (default is $HOME/.toolkit.yaml)")
	// rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// rootCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")
}
