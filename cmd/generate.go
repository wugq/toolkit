package cmd

import (
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate TYPE",
	Short: "Generate things like password, UUID and so on",
	Long:  `Generate things like password, UUID and so on.`,
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
