package handler

import (
	"net/http"

	"github.com/Wammero/PVZ-service/internal/service"

	"github.com/go-chi/chi"
)

type Handler struct {
	services *service.Service
}

func New(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) SetupRoutes(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
}
