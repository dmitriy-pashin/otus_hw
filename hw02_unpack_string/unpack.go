package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(inputString string) (string, error) {
	// Place your code here
	minRepeats := 2
	prevRunePositionDiff := 1
	var prevRune rune
	var resultString strings.Builder

	for pos, char := range inputString {
		if unicode.IsDigit(char) && pos == 0 {
			return "", ErrInvalidString
		}

		if pos == 0 {
			resultString.WriteRune(char)
			continue
		}

		if unicode.IsDigit(char) {
			repeatsCount, err := strconv.Atoi(string(char))

			if err != nil {
				return "", err
			}

			if repeatsCount < minRepeats {
				return "", ErrInvalidString
			}

			prevRune = rune(inputString[pos-prevRunePositionDiff])

			if unicode.IsDigit(prevRune) {
				return "", ErrInvalidString
			}

			repeatedChar := strings.Repeat(string(prevRune), repeatsCount-prevRunePositionDiff)
			resultString.WriteString(repeatedChar)
			continue
		}

		resultString.WriteRune(char)
	}

	return resultString.String(), nil
}
