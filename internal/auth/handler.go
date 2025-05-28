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
	AuthService *AuthService
}

type AuthHandler struct {
	*configs.Config
	AuthService *AuthService
}

func NewAuthHandler(router *http.ServeMux, deps *AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}

	routes := []string{
		"POST /auth/login",
		"POST /auth/register",
	}

	routeHandlers := []func() http.HandlerFunc{
		handler.login,
		handler.register,
	}

	for i, route := range routes {
		log.Printf("[Auth] - [Handler] - [INFO] route: %s", route)
		router.HandleFunc(route, routeHandlers[i]())
	}
}

func (h *AuthHandler) login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := req.HandleBody[LoginRequest](w, r)
		if err != nil {
			return
		}
		log.Printf("[Auth] - [Handler] - [INFO] login: %s", payload)

		data := LoginResponse{
			Token: "1234567890",
		}
		res.Json(w, data, http.StatusOK)
	}
}

func (h *AuthHandler) register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](w, r)
		if err != nil {
			return
		}
		user, err := h.AuthService.Register(body.Email, body.Username, body.Password)
		if err != nil {
			return
		}
		res.Json(w, user, http.StatusCreated)
	}
}
