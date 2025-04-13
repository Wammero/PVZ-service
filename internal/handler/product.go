package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Wammero/PVZ-service/internal/service"
	"github.com/Wammero/PVZ-service/pkg/responsemaker"
)

type productHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *productHandler {
	return &productHandler{service: service}
}

func (h *productHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Type  string `json:"type"`
		PVZID string `json:"pvzId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responsemaker.WriteJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	product, err := h.service.AddProduct(r.Context(), req.Type, req.PVZID)
	if err != nil {
		responsemaker.WriteJSONError(w, fmt.Sprintf("Ошибка при создании приёмки: %v", err), http.StatusBadRequest)
		return
	}

	responsemaker.WriteJSONResponse(w, product, http.StatusCreated)
}
