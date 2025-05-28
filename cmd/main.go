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

func App() (http.Handler, *configs.Config) {
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

	//Events Handlers
	//===============================================
	go statService.AddClick()
	//===============================================

	//Middlewares
	//===============================================
	chain := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)
	//===============================================

	return chain(router), conf
}

func main() {
	app, conf := App()
	server := http.Server{
		Addr:    ":" + conf.Server.Port,
		Handler: app,
	}

	log.Printf("[CMD] - [INFO] Starting server on port %s", conf.Server.Port)
	server.ListenAndServe()
}
