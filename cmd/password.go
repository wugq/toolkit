package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var UseUppercase bool
var UseLowercase bool
var UseNumber bool
var UseSymbol bool
var Size = 8

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

	passwordCmd.Flags().BoolVarP(&UseUppercase, "uppercase", "u", false, "use upper case")
	passwordCmd.Flags().BoolVarP(&UseLowercase, "lowercase", "l", false, "use lower case")
	passwordCmd.Flags().BoolVarP(&UseNumber, "number", "n", false, "use number")
	passwordCmd.Flags().BoolVarP(&UseSymbol, "symbol", "s", false, "use symbol")
}

func runPassword(args []string) {
	//fmt.Printf("Arguments : %v\n", args)
	//fmt.Printf("UseUppercase %v\n", UseUppercase)
	//fmt.Printf("UseLowercase %v\n", UseLowercase)
	//fmt.Printf("UseNumber %v\n", UseNumber)
	//fmt.Printf("UseSymbol %v\n", UseSymbol)

	if !UseUppercase && !UseLowercase && !UseNumber && !UseSymbol {
		UseUppercase = true
		UseLowercase = true
		UseNumber = true
		UseSymbol = true
	}

	if len(args) == 1 {
		i, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Please enter a correct number")
			os.Exit(2)
		}

		Size = i
	}

	var password = generatePassword(Size, UseLowercase, UseUppercase, UseSymbol, UseNumber)
	fmt.Println(password)

}

func makeAllCharSet(lowercaseFlag bool, uppercaseFlag bool, symbolFlag bool, numberFlag bool) string {
	var allCharsetBuilder strings.Builder
	if lowercaseFlag {
		allCharsetBuilder.WriteString(lowerCharSet)
	}
	if uppercaseFlag {
		allCharsetBuilder.WriteString(upperCharSet)
	}
	if symbolFlag {
		allCharsetBuilder.WriteString(symbolCharSet)
	}
	if numberFlag {
		allCharsetBuilder.WriteString(numberSet)
	}
	return allCharsetBuilder.String()
}

func generatePassword(length int, lowercaseFlag bool, uppercaseFlag bool, symbolFlag bool, numberFlag bool) string {
	var password strings.Builder
	var allCharset = makeAllCharSet(lowercaseFlag, uppercaseFlag, symbolFlag, numberFlag)

	if lowercaseFlag {
		random := rand.Intn(len(lowerCharSet))
		password.WriteString(string(lowerCharSet[random]))
	}
	if uppercaseFlag {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}
	if symbolFlag {
		random := rand.Intn(len(symbolCharSet))
		password.WriteString(string(symbolCharSet[random]))
	}
	if numberFlag {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	remainingLength := length - password.Len()
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
