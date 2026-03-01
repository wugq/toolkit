package passwordRunner

import (
	"math/rand"
	"strings"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano())) //nolint:gosec

type PasswordFlag struct {
	UseUppercase bool
	UseLowercase bool
	UseNumber    bool
	UseSymbol    bool
	Size         int
}

var (
	lowerCharSet  = "abcdefghijklmnopqrstuvwxyz"
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
		password.WriteString(string(lowerCharSet[rng.Intn(len(lowerCharSet))]))
	}
	if flag.UseUppercase {
		password.WriteString(string(upperCharSet[rng.Intn(len(upperCharSet))]))
	}
	if flag.UseSymbol {
		password.WriteString(string(symbolCharSet[rng.Intn(len(symbolCharSet))]))
	}
	if flag.UseNumber {
		password.WriteString(string(numberSet[rng.Intn(len(numberSet))]))
	}

	remainingLength := flag.Size - password.Len()
	for i := 0; i < remainingLength; i++ {
		password.WriteString(string(allCharset[rng.Intn(len(allCharset))]))
	}
	inRune := []rune(password.String())
	rng.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}
