package stat

import (
	"app/test/configs"
	"app/test/pkg/res"
	"net/http"
	"time"
)

const (
	GroupByDay   = "day"
	GroupByMonth = "month"
)

type StatHandler struct {
	StatRepository *StatRepository
}

type StatHandlerDeps struct {
	StatHandler *StatHandler
	Config      *configs.Config
}

func NewStatHandler(router *http.ServeMux, deps *StatHandlerDeps) {
	handler := &StatHandler{
		StatRepository: deps.StatHandler.StatRepository,
	}

	//Routes
	//===============================================
	publicRoutes := []string{
		"GET /stat",
	}
	//===============================================

	//Handlers
	//===============================================
	publicHandlers := []func() http.HandlerFunc{
		handler.GetAll,
	}
	//===============================================

	//Register routes
	//===============================================
	for i, route := range publicRoutes {
		router.HandleFunc(route, publicHandlers[i]())
	}
	//===============================================
}

func (h *StatHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		from, err := time.Parse(time.DateOnly, r.URL.Query().Get("from"))
		if err != nil {
			http.Error(w, "Invalid from", http.StatusBadRequest)
			return
		}

		to, err := time.Parse(time.DateOnly, r.URL.Query().Get("to"))
		if err != nil {
			http.Error(w, "Invalid to", http.StatusBadRequest)
			return
		}

		by := r.URL.Query().Get("by")
		if by != GroupByDay && by != GroupByMonth {
			http.Error(w, "Invalid by", http.StatusBadRequest)
			return
		}

		stats := h.StatRepository.GetStats(by, from, to)

		res.Json(w, stats, http.StatusOK)
	}
}
