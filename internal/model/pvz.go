package model

import "time"

type PVZWithReceptions struct {
	PVZ        PVZ                     `json:"pvz"`
	Receptions []ReceptionWithProducts `json:"receptions"`
}

type PVZ struct {
	ID               string    `json:"id"`
	RegistrationDate time.Time `json:"registrationDate"`
	City             string    `json:"city"`
}

type ReceptionWithProducts struct {
	Reception Reception `json:"reception"`
	Products  []Product `json:"products"`
}

type Reception struct {
	ID     string    `json:"id"`
	Date   time.Time `json:"dateTime"`
	PVZID  string    `json:"pvzId"`
	Status string    `json:"status"`
}

type Product struct {
	ID          string    `json:"id"`
	Date        time.Time `json:"dateTime"`
	Type        string    `json:"type"`
	ReceptionID string    `json:"receptionId"`
}
