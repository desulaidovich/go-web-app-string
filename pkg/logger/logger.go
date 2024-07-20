package logger

import (
	"net/http"
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	*zerolog.Logger
}

func New() *Logger {
	writer := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: zerolog.TimeFormatUnix,
	}

	logger := zerolog.
		New(writer).
		With().
		Timestamp().
		Logger()

	return &Logger{
		&logger,
	}
}

func (l *Logger) Init(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l.Debug().
			Str("FROM", r.RemoteAddr).
			Str("METHOD", r.Method).
			Str("URL", r.URL.String()).
			Send()
		h.ServeHTTP(w, r)
	})
}
