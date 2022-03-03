package passwordRunner

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
