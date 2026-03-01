package cmd

import (
	"fmt"
	"toolkit/runner/uuidRunner"

	"github.com/spf13/cobra"
)

var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generate a UUID string.",
	Long: `Generate a random UUID (version 4) and print it to stdout.

Example:
  toolkit generate uuid`,
	Run: func(cmd *cobra.Command, args []string) {
		runUUID()
	},
}

func init() {
	generateCmd.AddCommand(uuidCmd)
}

func runUUID() {
	fmt.Println(uuidRunner.Generate())
}
