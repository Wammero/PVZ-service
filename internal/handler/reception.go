package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Wammero/PVZ-service/internal/service"
	"github.com/Wammero/PVZ-service/pkg/responsemaker"
)

type receptionHandler struct {
	service service.ReceptionService
}

func NewReceptionHandler(service service.ReceptionService) *receptionHandler {
	return &receptionHandler{service: service}
}

func (h *receptionHandler) CreateReception(w http.ResponseWriter, r *http.Request) {
	var req struct {
		PVZID string `json:"pvzId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responsemaker.WriteJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	receptionId, dateTime, err := h.service.CreateReception(r.Context(), req.PVZID)
	if err != nil {
		responsemaker.WriteJSONError(w, fmt.Sprintf("Ошибка при создании приёмки: %v", err), http.StatusBadRequest)
		return
	}

	resp := struct {
		ID       string `json:"id"`
		DateTime string `json:"dateTime"`
		PVZID    string `json:"pvzId"`
		Status   string `json:"status"`
	}{
		ID:       receptionId,
		DateTime: dateTime,
		PVZID:    req.PVZID,
		Status:   "in_progress",
	}

	responsemaker.WriteJSONResponse(w, resp, http.StatusOK)
}
