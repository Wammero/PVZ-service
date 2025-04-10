package handler

import (
	"net/http"

	"github.com/Wammero/PVZ-service/internal/service"
)

type authHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) *authHandler {
	return &authHandler{service: service}
}

func (h *authHandler) Register(w http.ResponseWriter, r *http.Request) {

}

func (h *authHandler) Login(w http.ResponseWriter, r *http.Request) {

}

func (h *authHandler) DummyLogin(w http.ResponseWriter, r *http.Request) {

}
