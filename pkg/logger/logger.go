package logger

import (
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Logger struct {
	*zerolog.Logger
}

func New() *Logger {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()

	return &Logger{
		&logger,
	}
}

func (l *Logger) Init(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debug().
			Str("FROM", r.RemoteAddr).
			Str("METHOD", r.Method).
			Str("URL", r.URL.String()).
			Send()
		h.ServeHTTP(w, r)
	})
}
