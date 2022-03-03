package passwordRunner

import (
	"math/rand"
	"strings"
)

type PasswordFlag struct {
	UseUppercase bool
	UseLowercase bool
	UseNumber    bool
	UseSymbol    bool
	Size         int
}

var (
	lowerCharSet  = "abcdedfghijklmnopqrst"
	upperCharSet  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	symbolCharSet = "!@#$%&*"
	numberSet     = "0123456789"
)

func UpdateFlag(flag PasswordFlag) PasswordFlag {
	if !flag.UseUppercase && !flag.UseLowercase && !flag.UseNumber && !flag.UseSymbol {
		flag.UseUppercase = true
		flag.UseLowercase = true
		flag.UseNumber = true
		flag.UseSymbol = true
	}
	return flag
}

func MakeAllCharSet(flag PasswordFlag) string {
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

func GeneratePassword(flag PasswordFlag) string {
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
