package handler

import (
	"github.com/Wammero/PVZ-service/internal/service"
	"github.com/Wammero/PVZ-service/pkg/jwt"

	"github.com/go-chi/chi"
)

type Handler struct {
	services *service.Service
}

func New(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) SetupRoutes(r *chi.Mux) {
	r.Post("/dummyLogin", h.DummyLogin)
	r.Post("/register", h.Register)
	r.Post("/login", h.Login)

	r.Group(func(r chi.Router) {
		r.Use(jwt.JWTValidator)

		r.Route("/pvz", func(r chi.Router) {
			r.Post("/", h.CreatePVZ)
			r.Get("/", h.GetPVZList)
			r.Post("/{pvzId}/close_last_reception", h.CloseLastReception)
			r.Post("/{pvzId}/delete_last_product", h.DeleteLastProduct)
		})

		r.Post("/receptions", h.CreateReception)

		r.Post("/products", h.AddProduct)
	})
}
