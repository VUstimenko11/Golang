package hw02unpackstring

import (
	"errors"
	"strconv"
	"unicode"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	if input == "" {
        return "", nil
    }

    var result []rune
    var prev rune
    prevIsLetterOrNewline := false

    runes := []rune(input)

    for i := 0; i < len(runes); i++ {
        curr := runes[i]

        if unicode.IsLetter(curr) || curr == '\n' {
            result = append(result, curr)
            prev = curr
            prevIsLetterOrNewline = true
        } else if unicode.IsDigit(curr) {
            if i+1 < len(runes) && unicode.IsDigit(runes[i+1]) {
                return "", ErrInvalidString
            }
            if !prevIsLetterOrNewline {
                return "", ErrInvalidString
            }
            count, err := strconv.Atoi(string(curr))
            if err != nil {
                return "", ErrInvalidString
            }
            if count == 0 {
                result = result[:len(result)-1]
            } else {
                for j := 1; j < count; j++ {
                    result = append(result, prev)
                }
            }
        } else {
            result = append(result, curr)
            prevIsLetterOrNewline = false
        }
    }
    return string(result), nil
}

// функция, которая превращает символ \n в видимый "\n" в выводе
func escapeNewlines(s string) string {
	return strings.ReplaceAll(s, "\n", "\\n")
}
