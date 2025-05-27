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
	log.Printf("conf: %+v", conf)

	router := http.NewServeMux()
	auth.NewAuthHandler(router, &auth.AuthHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	log.Println("Starting server on port 3000")
	server.ListenAndServe()
}
