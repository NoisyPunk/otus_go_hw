package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	runeSlice := []rune(s)
	var result string

	if s != "" && unicode.IsDigit(runeSlice[0]) {
		return "", ErrInvalidString
	}

	for i, value := range runeSlice {
		switch {
		case unicode.IsLetter(value):
			result += string(value)
		case unicode.IsDigit(value) && unicode.IsDigit(runeSlice[i-1]):
			return "", ErrInvalidString
		case unicode.IsDigit(value) && !unicode.IsDigit(runeSlice[i-1]) && int(value-'0') > 1:
			result += strings.Repeat(string(runeSlice[i-1]), int((value-'0')-1))
		case unicode.IsDigit(value) && !unicode.IsDigit(runeSlice[i-1]) && int(value-'0') <= 1:
			result += string(value)
			result = strings.ReplaceAll(result, string(runeSlice[i-1:i+1]), "")
		}
	}
	return result, nil
}
