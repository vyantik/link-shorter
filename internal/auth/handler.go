package auth

import (
	"app/test/configs"
	"app/test/pkg/req"
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
		payload, err := req.HandleBody[LoginRequest](w, r)
		if err != nil {
			log.Printf("[ERROR] login: %s", err)
			return
		}
		log.Printf("[INFO] login: %s", payload)

		data := LoginResponse{
			Token: "1234567890",
		}
		res.Json(w, data, http.StatusOK)
	}
}

func (h *AuthHandler) register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("[INFO] register")
		data := map[string]string{"message": "register success"}
		res.Json(w, data, http.StatusCreated)
	}
}
