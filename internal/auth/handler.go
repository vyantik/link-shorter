package auth

import (
	"app/test/configs"
	"app/test/pkg/jwt"
	"app/test/pkg/req"
	"app/test/pkg/res"
	"log"
	"net/http"
)

type AuthHandler struct {
	*configs.Config
	AuthService *AuthService
}

type AuthHandlerDeps struct {
	AuthHandler *AuthHandler
}

func NewAuthHandler(router *http.ServeMux, deps *AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.AuthHandler.Config,
		AuthService: deps.AuthHandler.AuthService,
	}

	publicRoutes := []string{
		"POST /auth/login",
		"POST /auth/register",
	}

	publicHandlers := []func() http.HandlerFunc{
		handler.Login,
		handler.Register,
	}

	for i, route := range publicRoutes {
		log.Printf("[Auth] - [Handler] - [INFO] route: %s", route)
		router.HandleFunc(route, publicHandlers[i]())
	}
}

func (h *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](w, r)
		if err != nil {
			return
		}
		userEmail, err := h.AuthService.Login(body.Email, body.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
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

func (h *AuthHandler) Register() http.HandlerFunc {
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
