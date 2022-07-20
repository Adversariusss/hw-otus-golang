package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var (
	ErrFirstCharIsDigit = errors.New("first char is digit")
	ErrTwoDigits        = errors.New("two digits in a string")
)

func Unpack(str string) (string, error) {
	if str == "" {
		return "", nil
	}

	runeStr := []rune(str)

	if unicode.IsDigit(runeStr[0]) {
		return "", ErrFirstCharIsDigit
	}

	var result strings.Builder
	var digitCounter int

	for i, char := range runeStr {
		if unicode.IsDigit(char) {
			digitCounter++

			if digitCounter > 1 {
				return "", ErrTwoDigits
			}

			continue
		}

		if i == len(runeStr)-1 {
			result.WriteRune(char)
			continue
		}

		nextChar := runeStr[i+1]

		digit, err := strconv.Atoi(string(nextChar))
		if err != nil {
			result.WriteRune(char)
			continue
		}

		digitCounter = 0
		if digit <= 0 {
			continue
		}

		result.WriteString(strings.Repeat(string(char), digit))
	}
	return result.String(), nil
}
