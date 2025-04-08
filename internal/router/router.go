package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func New() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.URLFormat)
	router.Use(middleware.Logger)
	return router
}
