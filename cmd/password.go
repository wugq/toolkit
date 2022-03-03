package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"toolkit/runner/passwordRunner"
)

var passwordFlag passwordRunner.PasswordFlag

var passwordCmd = &cobra.Command{
	Use:   "password LENGTH",
	Short: "Generate random password.",
	Long:  `Generate random password.`,
	Run: func(cmd *cobra.Command, args []string) {
		runPassword()
	},
}

func init() {
	generateCmd.AddCommand(passwordCmd)

	passwordCmd.Flags().BoolVarP(&passwordFlag.UseUppercase, "uppercase", "u", false, "use upper case")
	passwordCmd.Flags().BoolVarP(&passwordFlag.UseLowercase, "lowercase", "l", false, "use lower case")
	passwordCmd.Flags().BoolVarP(&passwordFlag.UseNumber, "number", "n", false, "use number")
	passwordCmd.Flags().BoolVarP(&passwordFlag.UseSymbol, "symbol", "s", false, "use symbol")
	passwordCmd.Flags().IntVarP(&passwordFlag.Size, "length", "L", 8, "password length")
}

func runPassword() {
	passwordFlag = passwordRunner.UpdateFlag(passwordFlag)

	password := passwordRunner.GeneratePassword(passwordFlag)
	fmt.Println(password)
}
