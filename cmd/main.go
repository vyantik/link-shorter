package main

import (
	"app/test/configs"
	"app/test/internal/auth"
	"app/test/pkg/db"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)

	router := http.NewServeMux()

	auth.NewAuthHandler(router, &auth.AuthHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":" + conf.Server.Port,
		Handler: router,
	}

	log.Printf("[CMD] - [INFO] Starting server on port %s", conf.Server.Port)
	server.ListenAndServe()
}
