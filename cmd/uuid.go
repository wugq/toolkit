package cmd

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// uuidCmd represents the uuid command
var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generate a UUID string.",
	Long:  `Generate a UUID string.`,
	Run: func(cmd *cobra.Command, args []string) {
		runUUID()
	},
}

func init() {
	generateCmd.AddCommand(uuidCmd)
}

func runUUID() {
	id := uuid.New()
	fmt.Println(id.String())
}
