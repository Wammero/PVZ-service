package handler

import (
	"github.com/Wammero/PVZ-service/internal/middleware"
	"github.com/Wammero/PVZ-service/internal/service"

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
		r.Use(middleware.JWTValidator)

		r.Route("/pvz", func(r chi.Router) {
			r.With(middleware.RequireRole("moderator")).Post("/", h.PVZHandler.CreatePVZ)
			r.With(middleware.RequireRole("employee")).Get("/", h.PVZHandler.GetPVZList)
			r.With(middleware.RequireRole("employee")).Post("/{pvzId}/close_last_reception", h.PVZHandler.CloseLastReception)
			r.With(middleware.RequireRole("employee")).Post("/{pvzId}/delete_last_product", h.PVZHandler.DeleteLastProduct)
		})

		r.With(middleware.RequireRole("employee")).Post("/receptions", h.ReceptionHandler.CreateReception)

		r.With(middleware.RequireRole("employee")).Post("/products", h.ProductHandler.AddProduct)
	})
}
