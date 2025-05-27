package link

import (
	"net/http"
)

type LinkHandler struct {
}

func NewLinkHandler(router *http.ServeMux) {
	handler := &LinkHandler{}
	router.HandleFunc("GET /link", handler.GetLink())
	router.HandleFunc("POST /link", handler.CreateLink())
	router.HandleFunc("PUT /link", handler.UpdateLink())
	router.HandleFunc("DELETE /link", handler.DeleteLink())
}

func (h *LinkHandler) GetLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *LinkHandler) CreateLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *LinkHandler) UpdateLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *LinkHandler) DeleteLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
