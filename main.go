package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"unicode"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	loggingRequest := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
			log.Debug().
				Str("FROM", r.RemoteAddr).
				Str("METHOD", r.Method).
				Str("URL", r.URL.String()).
				Send()
			h.ServeHTTP(w, r)
		})
	}

	router := http.NewServeMux()

	// .../decrypt?value=VALUE
	router.HandleFunc("GET /decrypt", func(w http.ResponseWriter, r *http.Request) {
		value := r.URL.Query().Get("value")

		if len(value) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("NULL"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		reqJSON, _ := json.Marshal(map[string]string{
			"decrypt": DecryptLetters(value),
		})
		w.Write(reqJSON)
	})

	// .../encrypt?value=VALUE
	router.HandleFunc("GET /encrypt", func(w http.ResponseWriter, r *http.Request) {
		value := r.URL.Query().Get("value")

		if len(value) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("NULL"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		reqJSON, _ := json.Marshal(map[string]string{
			"encrypt": EncryptLetter(value),
		})
		w.Write(reqJSON)
	})

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: loggingRequest(router),
	}

	log.Info().
		Str(`Server run`, server.Addr).
		Send()

	err := server.ListenAndServe()
	log.Error().Msg(err.Error())

}

func EncryptLetter(input string) string {
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

func LetterPow(input string) string {
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

func ExpandExpression(input string) (string, bool) {
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

func DecryptLetters(input string) string {
	buff, ok := ExpandExpression(input)

	for {
		if !ok {
			break
		}

		buff, ok = ExpandExpression(buff)
	}

	buff = LetterPow(buff)

	return buff
}
