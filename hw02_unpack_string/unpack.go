package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const (
	BACKSLASH rune = 92
	ZERO      rune = 48
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	// Place your code here.

	r := []rune(s)

	lenStr := len(r)

	if lenStr == 0 {
		return "", nil
	}
	if unicode.IsDigit(r[0]) {
		return "", ErrInvalidString
	}

	var b strings.Builder

	i := 0
	for i < lenStr {
		switch {
		case !unicode.IsDigit(r[i]) && r[i] != BACKSLASH:
			if i < lenStr-1 && r[i+1] == ZERO { // aaa0b => aab
				i += 2
			} else { // write the symbol as it is
				b.WriteString(string(r[i]))
				i++
			}

		case unicode.IsDigit(r[i]):
			if i < lenStr-1 && unicode.IsDigit(r[i+1]) { // two digits in a row "45"
				return "", ErrInvalidString
			}
			n, err := strconv.Atoi(string(r[i]))
			if err != nil {
				fmt.Println("error strconv - ", err)
			} else {
				b.WriteString(strings.Repeat(string(r[i-1]), n-1)) // "a4bc2d5e" => "aaaabccddddde". a + aaa
				i++
			}

		default:

			switch {
			case i == lenStr-1:
				return "", ErrInvalidString // `qwe\`
			case unicode.IsDigit(r[i+1]): // all with slashes
				b.WriteString(string(r[i+1]))
				i += 2
			case r[i+1] == BACKSLASH: // more than one slash
				b.WriteString(string(r[i+1]))
				i += 2
			case r[i+1] != BACKSLASH: // `qwe\t`
				return "", ErrInvalidString
			}
		}
	}
	return b.String(), nil
}
