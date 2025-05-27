package main

import (
	"app/test/configs"
	"app/test/internal/auth"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	conf := configs.LoadConfig()

	router := http.NewServeMux()
	auth.NewAuthHandler(router, &auth.AuthHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":" + conf.Server.Port,
		Handler: router,
	}

	log.Printf("Starting server on port %s", conf.Server.Port)
	server.ListenAndServe()
}
