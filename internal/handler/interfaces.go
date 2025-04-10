package handler

import "net/http"

type AuthHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	DummyLogin(w http.ResponseWriter, r *http.Request)
}

type PVZHandler interface {
	CreatePVZ(w http.ResponseWriter, r *http.Request)
	GetPVZList(w http.ResponseWriter, r *http.Request)
	CloseLastReception(w http.ResponseWriter, r *http.Request)
	DeleteLastProduct(w http.ResponseWriter, r *http.Request)
}

type ReceptionHandler interface {
	CreateReception(w http.ResponseWriter, r *http.Request)
}

type ProductHandler interface {
	AddProduct(w http.ResponseWriter, r *http.Request)
}
