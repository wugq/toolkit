package cmd

import (
	"fmt"
	"os"
	"toolkit/runner/tokenrunner"

	"github.com/spf13/cobra"
)

type TokenCmdData struct {
	length int
}

var tokenCmdData TokenCmdData

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Generate a secure random hex token.",
	Long: `Generate a cryptographically secure random hex token.

  -L / --length  Number of random bytes (default: 16, produces a 32-char hex string)

Examples:
  toolkit generate token
  toolkit generate token -L 32`,
	Run: func(cmd *cobra.Command, args []string) {
		runToken()
	},
}

func init() {
	generateCmd.AddCommand(tokenCmd)
	tokenCmd.Flags().IntVarP(&tokenCmdData.length, "length", "L", 16, "Number of random bytes")
}

func runToken() {
	token, err := tokenrunner.Generate(tokenCmdData.length)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(token)
}
