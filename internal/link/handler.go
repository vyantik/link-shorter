package link

import (
	"app/test/pkg/req"
	"app/test/pkg/res"
	"log"
	"net/http"
)

type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
}

type LinkHandler struct {
	LinkRepository *LinkRepository
}

func NewLinkHandler(router *http.ServeMux, deps *LinkHandlerDeps) {
	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
	}

	routes := []string{
		"GET /{hash}",
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
		hash := r.PathValue("hash")
		link, err := h.LinkRepository.GetByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}
}

func (h *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LinkCreateRequest](w, r)
		if err != nil {
			return
		}

		link := NewLink(body.Url)
		link, err = h.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Json(w, link, http.StatusCreated)

		log.Printf("[Link] - [Handler] - [INFO] create: %s", body.Url)

	}
}

func (h *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		log.Printf("[Link] - [Handler] - [INFO] delete: %s", id)
	}
}
