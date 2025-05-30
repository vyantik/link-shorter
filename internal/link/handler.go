package link

import (
	"app/test/configs"
	"app/test/pkg/event"
	"app/test/pkg/middleware"
	"app/test/pkg/req"
	"app/test/pkg/res"
	"log"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type LinkHandler struct {
	LinkRepository *LinkRepository
	EventBus       *event.EventBus
}

type LinkHandlerDeps struct {
	LinkHandler *LinkHandler
	Config      *configs.Config
}

func NewLinkHandler(router *http.ServeMux, deps *LinkHandlerDeps) {
	handler := &LinkHandler{
		LinkRepository: deps.LinkHandler.LinkRepository,
		EventBus:       deps.LinkHandler.EventBus,
	}

	//Routes
	//===============================================
	publicRoutes := []string{
		"GET /{hash}",
	}
	privateRoutes := []string{
		"GET /link",
		"POST /link",
		"PATCH /link/{id}",
		"DELETE /link/{id}",
	}
	//===============================================

	//Handlers
	//===============================================
	publicHandlers := []func() http.HandlerFunc{
		handler.GoTo,
	}
	privateHandlers := []func() http.HandlerFunc{
		handler.GetAll,
		handler.Create,
		handler.Update,
		handler.Delete,
	}
	//===============================================

	//Register routes
	//===============================================
	for i, route := range privateRoutes {
		log.Printf("[Link] - [Handler] - [INFO] route: %s", route)
		router.Handle(route, middleware.IsAuthed(privateHandlers[i](), deps.Config))
	}
	for i, route := range publicRoutes {
		log.Printf("[Link] - [Handler] - [INFO] route: %s", route)
		router.Handle(route, publicHandlers[i]())
	}
	//===============================================
}

func (h *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		link, err := h.LinkRepository.GetByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		go h.EventBus.Publish(event.Event{
			Type: event.LinkVisited,
			Data: link.ID,
		})
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

		for {
			existedLink, _ := h.LinkRepository.GetByHash(link.Hash)
			if existedLink == nil {
				break
			}
			link.GenerateHash()
		}

		link, err = h.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Json(w, link, http.StatusCreated)
	}
}

func (h *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LinkUpdateRequest](w, r)
		if err != nil {
			return
		}

		email, ok := r.Context().Value(middleware.ContextEmailKey).(string)
		if ok {
			log.Println(email)
		}

		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		link, err := h.LinkRepository.Update(&Link{
			Model: gorm.Model{ID: uint(id)},
			Url:   body.Url,
			Hash:  body.Hash,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, link, http.StatusOK)
	}
}

func (h *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = h.LinkRepository.GetById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		err = h.LinkRepository.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Json(w, nil, http.StatusNoContent)
	}
}

func (h *LinkHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			http.Error(w, "Invalid limit", http.StatusBadRequest)
			return
		}

		offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err != nil {
			http.Error(w, "Invalid offset", http.StatusBadRequest)
			return
		}

		links := h.LinkRepository.GetAll(limit, offset)
		count := h.LinkRepository.GetCount()
		res.Json(w, LinkGetAllResponse{
			Links: links,
			Count: count,
		}, http.StatusOK)
	}
}
