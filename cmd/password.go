package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"math/rand"
	"os"
	"strings"
	"toolkit/runner/passwordRunner"
)

var passwordFlag passwordRunner.PasswordFlag

var (
	lowerCharSet  = "abcdedfghijklmnopqrst"
	upperCharSet  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	symbolCharSet = "!@#$%&*"
	numberSet     = "0123456789"
)

var passwordCmd = &cobra.Command{
	Use:   "password LENGTH",
	Short: "Generate random password.",
	Long:  `Generate random password.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			err := cmd.Help()
			if err != nil {
				return
			}
			os.Exit(0)
		}
		runPassword(args)
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

func runPassword(args []string) {
	passwordFlag = passwordRunner.UpdateFlag(passwordFlag)

	fmt.Println(GeneratePassword(passwordFlag))

}

func MakeAllCharSet(flag passwordRunner.PasswordFlag) string {
	var allCharsetBuilder strings.Builder
	if flag.UseLowercase {
		allCharsetBuilder.WriteString(lowerCharSet)
	}
	if flag.UseUppercase {
		allCharsetBuilder.WriteString(upperCharSet)
	}
	if flag.UseSymbol {
		allCharsetBuilder.WriteString(symbolCharSet)
	}
	if flag.UseNumber {
		allCharsetBuilder.WriteString(numberSet)
	}
	return allCharsetBuilder.String()
}

func GeneratePassword(flag passwordRunner.PasswordFlag) string {
	var password strings.Builder
	var allCharset = MakeAllCharSet(flag)

	if flag.UseLowercase {
		random := rand.Intn(len(lowerCharSet))
		password.WriteString(string(lowerCharSet[random]))
	}
	if flag.UseUppercase {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}
	if flag.UseSymbol {
		random := rand.Intn(len(symbolCharSet))
		password.WriteString(string(symbolCharSet[random]))
	}
	if flag.UseNumber {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	remainingLength := flag.Size - password.Len()
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharset))
		password.WriteString(string(allCharset[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}
