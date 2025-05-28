package main

import (
	"app/test/configs"
	"app/test/internal/auth"
	"app/test/internal/link"
	"app/test/internal/stat"
	"app/test/internal/user"
	"app/test/pkg/db"
	"app/test/pkg/event"
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

	//EventBus
	//===============================================
	eventBus := event.NewEventBus()
	//===============================================

	//Repositories
	//===============================================
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	statRepository := stat.NewStatRepository(db)
	//===============================================

	//Services
	//===============================================
	authService := auth.NewAuthService(userRepository)
	statService := stat.NewStatService(&stat.StatServiceDeps{
		StatService: &stat.StatService{
			EventBus:       eventBus,
			StatRepository: statRepository,
		},
	})
	//===============================================

	//Handlers
	//===============================================
	auth.NewAuthHandler(router, &auth.AuthHandlerDeps{
		AuthHandler: &auth.AuthHandler{
			Config:      conf,
			AuthService: authService,
		},
	})
	link.NewLinkHandler(router, &link.LinkHandlerDeps{
		Config: conf,
		LinkHandler: &link.LinkHandler{
			LinkRepository: linkRepository,
			EventBus:       eventBus,
		},
	})
	stat.NewStatHandler(router, &stat.StatHandlerDeps{
		StatHandler: &stat.StatHandler{
			StatRepository: statRepository,
		},
		Config: conf,
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

	//Events Handlers
	//===============================================
	go statService.AddClick()
	//===============================================

	log.Printf("[CMD] - [INFO] Starting server on port %s", conf.Server.Port)
	server.ListenAndServe()
}
