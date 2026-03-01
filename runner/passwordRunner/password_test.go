package passwordrunner

import (
	"strings"
	"testing"
)

func TestUpdateFlag_AllDefault(t *testing.T) {
	flag := UpdateFlag(PasswordFlag{Size: 12})
	if !flag.UseUppercase || !flag.UseLowercase || !flag.UseNumber || !flag.UseSymbol {
		t.Error("expected all flags enabled when none set")
	}
}

func TestUpdateFlag_PartialFlags_NotOverridden(t *testing.T) {
	flag := UpdateFlag(PasswordFlag{UseUppercase: true, Size: 12})
	if !flag.UseUppercase {
		t.Error("UseUppercase should remain true")
	}
	if flag.UseLowercase || flag.UseNumber || flag.UseSymbol {
		t.Error("other flags should remain false when at least one is set")
	}
}

func TestUpdateFlag_AllSet_Unchanged(t *testing.T) {
	input := PasswordFlag{UseUppercase: true, UseLowercase: true, UseNumber: true, UseSymbol: true, Size: 8}
	got := UpdateFlag(input)
	if got != input {
		t.Error("all-set flag should be returned unchanged")
	}
}

func TestMakeAllCharSet_Lower(t *testing.T) {
	cs := MakeAllCharSet(PasswordFlag{UseLowercase: true})
	if cs != lowerCharSet {
		t.Errorf("got %q, want %q", cs, lowerCharSet)
	}
}

func TestMakeAllCharSet_Upper(t *testing.T) {
	cs := MakeAllCharSet(PasswordFlag{UseUppercase: true})
	if cs != upperCharSet {
		t.Errorf("got %q, want %q", cs, upperCharSet)
	}
}

func TestMakeAllCharSet_Combined(t *testing.T) {
	cs := MakeAllCharSet(PasswordFlag{UseLowercase: true, UseUppercase: true, UseNumber: true, UseSymbol: true})
	for _, r := range lowerCharSet + upperCharSet + numberSet + symbolCharSet {
		if !strings.ContainsRune(cs, r) {
			t.Errorf("charset missing expected char %q", r)
		}
	}
}

func TestMakeAllCharSet_Empty(t *testing.T) {
	cs := MakeAllCharSet(PasswordFlag{})
	if cs != "" {
		t.Errorf("expected empty charset, got %q", cs)
	}
}

func TestGeneratePassword_CorrectLength(t *testing.T) {
	for _, size := range []int{8, 12, 16, 32} {
		flag := PasswordFlag{UseUppercase: true, UseLowercase: true, UseNumber: true, UseSymbol: true, Size: size}
		pw := GeneratePassword(flag)
		if len(pw) != size {
			t.Errorf("size %d: got length %d", size, len(pw))
		}
	}
}

func TestGeneratePassword_ContainsEachCharSet(t *testing.T) {
	flag := PasswordFlag{UseUppercase: true, UseLowercase: true, UseNumber: true, UseSymbol: true, Size: 20}
	pw := GeneratePassword(flag)
	if !strings.ContainsAny(pw, lowerCharSet) {
		t.Errorf("password %q missing lowercase characters", pw)
	}
	if !strings.ContainsAny(pw, upperCharSet) {
		t.Errorf("password %q missing uppercase characters", pw)
	}
	if !strings.ContainsAny(pw, numberSet) {
		t.Errorf("password %q missing digit characters", pw)
	}
	if !strings.ContainsAny(pw, symbolCharSet) {
		t.Errorf("password %q missing symbol characters", pw)
	}
}

func TestGeneratePassword_OnlyLowercase(t *testing.T) {
	flag := PasswordFlag{UseLowercase: true, Size: 10}
	pw := GeneratePassword(flag)
	if len(pw) != 10 {
		t.Errorf("expected length 10, got %d", len(pw))
	}
	for _, r := range pw {
		if !strings.ContainsRune(lowerCharSet, r) {
			t.Errorf("password %q contains non-lowercase char %q", pw, r)
		}
	}
}

func TestGeneratePassword_Randomness(t *testing.T) {
	flag := PasswordFlag{UseUppercase: true, UseLowercase: true, UseNumber: true, UseSymbol: true, Size: 16}
	pw1 := GeneratePassword(flag)
	pw2 := GeneratePassword(flag)
	if pw1 == pw2 {
		t.Error("two generated passwords should not be equal (extremely unlikely)")
	}
}
