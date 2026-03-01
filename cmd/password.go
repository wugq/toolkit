package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"toolkit/runner/password"
)

var passwordFlag password.PasswordFlag

var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "Generate random password.",
	Long: `Generate a random password.

By default uses all character sets (uppercase, lowercase, numbers, symbols).
Specify one or more flags to restrict to those sets only:
  -u / --uppercase  Include uppercase letters (A-Z)
  -l / --lowercase  Include lowercase letters (a-z)
  -n / --number     Include digits (0-9)
  -s / --symbol     Include symbols
  -L / --length     Password length (default: 8)

Examples:
  toolkit generate password
  toolkit generate password -u -l -n -L 16
  toolkit generate password -n -s -L 12`,
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
	passwordFlag = password.UpdateFlag(passwordFlag)

	password := password.GeneratePassword(passwordFlag)
	fmt.Println(password)
}
