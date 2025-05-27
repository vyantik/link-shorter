package link

import (
	"app/test/configs"
	"log"
	"net/http"
)

type LinkHandlerDeps struct {
	*configs.Config
}

type LinkHandler struct {
	*configs.Config
}

func NewLinkHandler(router *http.ServeMux, deps *LinkHandlerDeps) {
	handler := &LinkHandler{
		Config: deps.Config,
	}

	routes := []string{
		"GET /{alias}",
		"POST /link",
		"PATCH /link/{id}",
		"DELETE /link/{id}",
	}

	routeHandlers := []func() http.HandlerFunc{
		handler.GoTo,
		handler.Create,
		handler.Update,
		handler.Delete,
	}

	for i, route := range routes {
		log.Printf("[Link] - [Handler] - [INFO] route: %s", route)
		router.HandleFunc(route, routeHandlers[i]())
	}
}

func (h *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
