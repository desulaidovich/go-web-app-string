package cryproher

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

type Cryproher struct{}

func (c *Cryproher) EncryptLetters(input string) string {
	buff := encryptLetter(input)
	buff, ok := encryptExpr(buff)

	for {
		if !ok {
			break
		}
		buff, ok = encryptExpr(buff)
	}
	return buff
}

func encryptExpr(input string) (string, bool) {
	count := 1
	for i := 0; i < len(input); i++ {
		for j := i; j < len(input); j++ {
			inputLen := len(input[i:j])
			if inputLen == 0 {
				continue
			}
			if inputLen+j > len(input) {
				break
			}
			if input[i:j] == input[j:inputLen+j] {
				count++
				str := fmt.Sprintf("%s%d(%s)%s",
					input[:i], count, input[i:j], input[inputLen+j:])

				if strings.Index(str, "()") > 0 {
					return input, false
				}

				return str, true
			}
		}
	}
	return input, false
}

func encryptLetter(input string) string {
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
