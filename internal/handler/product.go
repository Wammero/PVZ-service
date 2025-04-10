package handler

import (
	"net/http"

	"github.com/Wammero/PVZ-service/internal/service"
)

type productHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *productHandler {
	return &productHandler{service: service}
}

func (h *productHandler) AddProduct(w http.ResponseWriter, r *http.Request) {

}
