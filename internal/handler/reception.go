package handler

import (
	"net/http"

	"github.com/Wammero/PVZ-service/internal/service"
)

type receptionHandler struct {
	service service.ReceptionService
}

func NewReceptionHandler(service service.ReceptionService) *receptionHandler {
	return &receptionHandler{service: service}
}

func (h *receptionHandler) CreateReception(w http.ResponseWriter, r *http.Request) {

}
