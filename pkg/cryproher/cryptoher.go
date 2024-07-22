package cryproher

import (
	"bytes"
	"strings"
	"unicode"
)

type Cryproher struct{}

func (c *Cryproher) EncryptExpr(input string) string {
	for i := range input {
		for j := range input {
			if i >= j {
				continue
			}
			firstSlice := input[i : j+1]
			secondSlice := input[j+1:]

			if firstSlice == secondSlice {
				return "2(" + firstSlice + ")"
			}

			if strings.Index(secondSlice, firstSlice) == 0 {
				return "2(" + input[i:j+1] + ")" + input[j+1+(len(input[i:j+1])-1)+1:]
			}
		}
	}
	return input
}

func (c *Cryproher) EncryptLetter(input string) string {
	inputeLen := len(input)
	buff := &bytes.Buffer{}
	char := byte(0)
	count := 1

	for i := 1; i < inputeLen; i++ {
		if input[i] == input[i-1] {
			count += 1
			char = input[i]
		} else {
			if count > 1 {
				buff.WriteRune(rune(count + '0'))
				buff.WriteByte(char)
			} else {
				buff.WriteByte(input[i-1])
			}
			count = 1
		}
	}

	if count > 1 {
		buff.WriteRune(rune(count + '0'))
		buff.WriteByte(char)
		return buff.String()
	}

	buff.WriteByte(input[inputeLen-1])
	return buff.String()
}

func letterPow(input string) string {
	buff := []rune{}

	for pos, char := range input {
		if unicode.IsLetter(char) {
			buff = append(buff, char)
		} else if unicode.IsDigit(char) {
			step := int(char - '0')
			for j := 0; j < step-1; j++ {
				buff = append(buff, rune(input[pos+1]))
			}
		}
	}

	return string(buff)
}

func expandExpression(input string) (string, bool) {
	buff := []rune{}

	lastChar := strings.Index(input, ")")
	firstChar := strings.LastIndex(input, "(")

	if lastChar == -1 || firstChar == -1 {
		return input, false
	}

	char := rune(input[firstChar-1])
	step := int(char - '0')
	offset := 0

	// Если перед скобочкой не будет числа,
	// то будем считать, что перед скобочкой была единица,
	// но так как символа единицы нет, то мы сместим все на 1
	if !unicode.IsDigit(char) {
		step = 1
		offset = 1
	}

	for i := 0; i < firstChar-1+offset; i++ {
		buff = append(buff, rune(input[i]))
	}

	for i := 0; i < step; i++ {
		for j := firstChar + 1; j < lastChar; j++ {
			buff = append(buff, rune(input[j]))
		}
	}

	for i := lastChar + 1; i < len(input); i++ {
		buff = append(buff, rune(input[i]))
	}

	return string(buff), true
}

func (c *Cryproher) DecryptLetters(input string) string {
	buff, ok := expandExpression(input)

	for {
		if !ok {
			break
		}

		buff, ok = expandExpression(buff)
	}

	buff = letterPow(buff)

	return buff
}
