package auth

import (
	"app/test/configs"
	"app/test/pkg/res"
	"log"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps *AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.login())
	router.HandleFunc("POST /auth/register", handler.register())
}

func (h *AuthHandler) login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("login: %s", h.Config.Auth.Secret)
		data := LoginResponse{
			Token: "1234567890",
		}
		res.Json(w, data, http.StatusOK)
	}
}

func (h *AuthHandler) register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("register")
		data := map[string]string{"message": "register success"}
		res.Json(w, data, http.StatusCreated)
	}
}
