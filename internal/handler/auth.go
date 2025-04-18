package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Wammero/PVZ-service/internal/model"
	"github.com/Wammero/PVZ-service/internal/service"
	"github.com/Wammero/PVZ-service/pkg/responsemaker"
)

type authHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) *authHandler {
	return &authHandler{service: service}
}

func (h *authHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		responsemaker.WriteJSONError(w, fmt.Sprintf("Failed to login: %v", err), http.StatusBadRequest)
		return
	}

	err := h.service.Register(r.Context(), user.Email, user.Password, user.Role)
	if err != nil {
		responsemaker.WriteJSONError(w, fmt.Sprintf("Failed to login: %v", err), http.StatusInternalServerError)
		return
	}

	responsemaker.WriteJSONResponse(w, user, http.StatusCreated)
}

func (h *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user model.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		responsemaker.WriteJSONError(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	token, err := h.service.Login(r.Context(), user.Email, user.Password)
	if err != nil {
		responsemaker.WriteJSONError(w, fmt.Sprintf("Failed to login: %v", err), http.StatusInternalServerError)
		return
	}

	responsemaker.WriteJSONResponse(w, token, http.StatusOK)
}

func (h *authHandler) DummyLogin(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Role string `json:"role"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responsemaker.WriteJSONError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	tokenstring, err := h.service.DummyLogin(r.Context(), req.Role)
	if err != nil {
		responsemaker.WriteJSONError(w, "Failed to login", http.StatusBadRequest)
		return
	}

	responsemaker.WriteJSONResponse(w, tokenstring, http.StatusOK)
}
