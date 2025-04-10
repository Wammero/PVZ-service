package handler

import (
	"net/http"

	"github.com/Wammero/PVZ-service/internal/service"
)

type pVZHandler struct {
	service service.PVZService
}

func NewPVZHandler(service service.PVZService) *pVZHandler {
	return &pVZHandler{service: service}
}

func (h *pVZHandler) CreatePVZ(w http.ResponseWriter, r *http.Request) {

}

func (h *pVZHandler) GetPVZList(w http.ResponseWriter, r *http.Request) {

}

func (h *pVZHandler) CloseLastReception(w http.ResponseWriter, r *http.Request) {

}

func (h *pVZHandler) DeleteLastProduct(w http.ResponseWriter, r *http.Request) {

}
