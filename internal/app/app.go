package app

import (
	"app/internal/routes"
	"app/pkg/logger"
	"net/http"
)

func Run() {
	log := logger.New()

	router := http.NewServeMux()

	router.HandleFunc("GET /encrypt", routes.Encrypt)
	router.HandleFunc("GET /decrypt", routes.Decrypt)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: log.Init(router),
	}

	log.Info().Str(`Server run`, server.Addr).Send()
	log.Error().Msg(server.ListenAndServe().Error())
}
