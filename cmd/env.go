package cmd

import (
	"fmt"
	"toolkit/runner/env"

	"github.com/spf13/cobra"
)

type EnvCmdData struct {
	pretty bool
}

var envCmdData EnvCmdData

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "List environment variables.",
	Long: `List all environment variables.

Use -p / --pretty to format PATH-like variables on multiple lines.

Examples:
  toolkit env
  toolkit env -p`,
	Run: func(cmd *cobra.Command, args []string) {
		runEnv()
	},
}

func init() {
	rootCmd.AddCommand(envCmd)
	envCmd.Flags().BoolVarP(&envCmdData.pretty, "pretty", "p", false, "Pretty-print PATH-like variables")
}

func runEnv() {
	for _, line := range env.Env(envCmdData.pretty) {
		fmt.Println(line)
	}
}
