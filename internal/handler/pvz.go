package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	var req struct {
		ID               string `json:"id"`
		RegistrationDate string `json:"registrationDate"`
		City             string `json:"city"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responsemaker.WriteJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.service.CreatePVZ(r.Context(), req.ID, req.RegistrationDate, req.City)
	if err != nil {
		responsemaker.WriteJSONError(w, fmt.Sprintf("Failed to create PVZ: %v", err), http.StatusInternalServerError)
		return
	}

	resp := struct {
		ID               string `json:"id"`
		RegistrationDate string `json:"registrationDate"`
		City             string `json:"city"`
	}{
		ID:               req.ID,
		RegistrationDate: req.RegistrationDate,
		City:             req.City,
	}

	responsemaker.WriteJSONResponse(w, resp, http.StatusOK)
}

func (h *pVZHandler) GetPVZList(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	startDateStr := query.Get("startDate")
	endDateStr := query.Get("endDate")
	pageStr := query.Get("page")
	limitStr := query.Get("limit")

	pvzList, err := h.service.GetPVZList(r.Context(), startDateStr, endDateStr, pageStr, limitStr)
	if err != nil {
		responsemaker.WriteJSONError(w, "не удалось получить данные о pvz: "+err.Error(), http.StatusBadRequest)
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

	err := h.service.CloseLastReception(r.Context(), pvzID)
	if err != nil {
		responsemaker.WriteJSONError(w, "не удалось закрыть приёмку: "+err.Error(), http.StatusBadRequest)
		return
	}

	responsemaker.WriteJSONResponse(w, "", http.StatusOK)
}

func (h *pVZHandler) DeleteLastProduct(w http.ResponseWriter, r *http.Request) {
	pvzID := chi.URLParam(r, "pvzId")

	if pvzID == "" {
		responsemaker.WriteJSONError(w, "missing pvz_id in URL", http.StatusBadRequest)
		return
	}

	err := h.service.DeleteLastProduct(r.Context(), pvzID)
	if err != nil {
		responsemaker.WriteJSONError(w, "не удалось закрыть приёмку: "+err.Error(), http.StatusBadRequest)
		return
	}

	responsemaker.WriteJSONResponse(w, "", http.StatusOK)

}
