package auth

import (
	"app/test/configs"
	"app/test/pkg/jwt"
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

	publicRoutes := []string{
		"POST /auth/login",
		"POST /auth/register",
	}

	publicHandlers := []func() http.HandlerFunc{
		handler.login,
		handler.register,
	}

	for i, route := range publicRoutes {
		log.Printf("[Auth] - [Handler] - [INFO] route: %s", route)
		router.HandleFunc(route, publicHandlers[i]())
	}
}

func (h *AuthHandler) login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](w, r)
		if err != nil {
			return
		}
		userEmail, err := h.AuthService.Login(body.Email, body.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		token, err := jwt.NewJWT(h.Config.Auth.Secret).Generate(jwt.JWTData{Email: userEmail})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Json(w, LoginResponse{Token: token}, http.StatusOK)
	}
}

func (h *AuthHandler) register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](w, r)
		if err != nil {
			return
		}
		userEmail, err := h.AuthService.Register(body.Email, body.Username, body.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		token, err := jwt.NewJWT(h.Config.Auth.Secret).Generate(jwt.JWTData{Email: userEmail})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Json(w, RegisterResponse{Token: token}, http.StatusCreated)
	}
}
