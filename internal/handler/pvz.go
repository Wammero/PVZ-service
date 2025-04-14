package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Wammero/PVZ-service/internal/metrics"
	"github.com/Wammero/PVZ-service/internal/model"
	"github.com/Wammero/PVZ-service/internal/service"
	"github.com/Wammero/PVZ-service/pkg/responsemaker"
	"github.com/go-chi/chi"
)

type pVZHandler struct {
	service service.PVZService
}

func NewPVZHandler(service service.PVZService) *pVZHandler {
	return &pVZHandler{service: service}
}

func (h *pVZHandler) CreatePVZ(w http.ResponseWriter, r *http.Request) {
	var pvz model.PVZ

	if err := json.NewDecoder(r.Body).Decode(&pvz); err != nil {
		responsemaker.WriteJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.service.CreatePVZ(r.Context(), pvz.ID, pvz.City, pvz.RegistrationDate)
	if err != nil {
		responsemaker.WriteJSONError(w, fmt.Sprintf("Failed to create PVZ: %v", err), http.StatusBadRequest)
		return
	}

	metrics.CreatedPVZ.Inc()

	responsemaker.WriteJSONResponse(w, pvz, http.StatusCreated)
}

func (h *pVZHandler) GetPVZList(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	startDateStr := query.Get("startDate")
	endDateStr := query.Get("endDate")
	pageStr := query.Get("page")
	limitStr := query.Get("limit")

	pvzList, err := h.service.GetPVZList(r.Context(), startDateStr, endDateStr, pageStr, limitStr)
	if err != nil {
		responsemaker.WriteJSONError(w, fmt.Sprintf("не удалось получить данные о pvz: %v", err), http.StatusBadRequest)
		return
	}

	responsemaker.WriteJSONResponse(w, pvzList, http.StatusOK)
}

func (h *pVZHandler) CloseLastReception(w http.ResponseWriter, r *http.Request) {
	pvzID := chi.URLParam(r, "pvzId")

	if pvzID == "" {
		responsemaker.WriteJSONError(w, "missing pvz_id in URL", http.StatusBadRequest)
		return
	}

	reception, err := h.service.CloseLastReception(r.Context(), pvzID)
	if err != nil {
		responsemaker.WriteJSONError(w, fmt.Sprintf("не удалось закрыть приёмку: %v", err), http.StatusBadRequest)
		return
	}

	responsemaker.WriteJSONResponse(w, reception, http.StatusOK)
}

func (h *pVZHandler) DeleteLastProduct(w http.ResponseWriter, r *http.Request) {
	pvzID := chi.URLParam(r, "pvzId")

	if pvzID == "" {
		responsemaker.WriteJSONError(w, "missing pvz_id in URL", http.StatusBadRequest)
		return
	}

	err := h.service.DeleteLastProduct(r.Context(), pvzID)
	if err != nil {
		responsemaker.WriteJSONError(w, fmt.Sprintf("не удалось закрыть приёмку: %v", err), http.StatusBadRequest)
		return
	}

	responsemaker.WriteJSONResponse(w, "", http.StatusOK)

}
