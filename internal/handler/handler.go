package handler

import (
	"github.com/Wammero/PVZ-service/internal/service"
	"github.com/Wammero/PVZ-service/pkg/jwt"

	"github.com/go-chi/chi"
)

type Handler struct {
	AuthHandler      AuthHandler
	PVZHandler       PVZHandler
	ReceptionHandler ReceptionHandler
	ProductHandler   ProductHandler
}

func New(services *service.Service) *Handler {
	return &Handler{
		AuthHandler:      NewAuthHandler(services.AuthService),
		PVZHandler:       NewPVZHandler(services.PVZService),
		ReceptionHandler: NewReceptionHandler(services.ReceptionService),
		ProductHandler:   NewProductHandler(services.ProductService),
	}
}

func (h *Handler) SetupRoutes(r *chi.Mux) {
	r.Post("/dummyLogin", h.AuthHandler.DummyLogin)
	r.Post("/register", h.AuthHandler.Register)
	r.Post("/login", h.AuthHandler.Login)

	r.Group(func(r chi.Router) {
		r.Use(jwt.JWTValidator)

		r.Route("/pvz", func(r chi.Router) {
			r.Post("/", h.PVZHandler.CreatePVZ)
			r.Get("/", h.PVZHandler.GetPVZList)
			r.Post("/{pvzId}/close_last_reception", h.PVZHandler.CloseLastReception)
			r.Post("/{pvzId}/delete_last_product", h.PVZHandler.DeleteLastProduct)
		})

		r.Post("/receptions", h.ReceptionHandler.CreateReception)

		r.Post("/products", h.ProductHandler.AddProduct)
	})
}
