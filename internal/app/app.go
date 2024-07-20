package app

import (
	"app/internal/routes"
	"app/pkg/logger"
	"net/http"
)

func Run() {
	log := logger.New()

	mux := http.NewServeMux()
	r := routes.New(mux)

	mux.HandleFunc("GET /encrypt", r.Encrypt)
	mux.HandleFunc("GET /decrypt", r.Decrypt)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: log.Init(mux),
	}

	log.Info().Str(`Server run`, server.Addr).Send()
	log.Error().Msg(server.ListenAndServe().Error())
}
