package main

import (
	"app/test/configs"
	"app/test/internal/auth"
	"app/test/internal/link"
	"app/test/pkg/db"
	"app/test/pkg/middleware"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)

	router := http.NewServeMux()

	//Repositories
	//===============================================
	linkRepository := link.NewLinkRepository(db)
	//===============================================

	//Handlers
	//===============================================
	auth.NewAuthHandler(router, &auth.AuthHandlerDeps{
		Config: conf,
	})
	link.NewLinkHandler(router, &link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})
	//===============================================

	//Middlewares
	//===============================================
	chain := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)
	//===============================================

	server := http.Server{
		Addr:    ":" + conf.Server.Port,
		Handler: chain(router),
	}

	log.Printf("[CMD] - [INFO] Starting server on port %s", conf.Server.Port)
	server.ListenAndServe()
}
