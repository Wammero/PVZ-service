package router

import (
	"github.com/Wammero/PVZ-service/internal/metrics"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func New() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.URLFormat)
	router.Use(middleware.Logger)

	router.Use(metrics.Middleware)

	return router
}
