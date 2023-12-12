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
	for i, value := range runeSlice {
		if unicode.IsDigit(value) && i == 0 {
			return "", ErrInvalidString
		}
		if unicode.IsDigit(value) && unicode.IsDigit(runeSlice[i-1]) {
			return "", ErrInvalidString
		}
		if unicode.IsLetter(value) {
			result += string(value)
			continue
		}
		if unicode.IsDigit(value) && !unicode.IsDigit(runeSlice[i-1]) && int(value-'0') > 1 {
			result += strings.Repeat(string(runeSlice[i-1]), int((value-'0')-1))
		}
		if unicode.IsDigit(value) && !unicode.IsDigit(runeSlice[i-1]) && int(value-'0') <= 1 {
			result += string(value)
			result = strings.ReplaceAll(result, string(runeSlice[i-1:i+1]), "")
		}
	}
	return result, nil
}
