package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"
)

const (
	baseURL     = "http://localhost:8081"
	pvzEndpoint = "/pvz"
	loginURL    = baseURL + "/dummyLogin"
)

var (
	apiPvzURL      = baseURL + pvzEndpoint
	apiReceptions  = baseURL + "/receptions"
	apiProducts    = baseURL + "/products"
	testPvzID      = "4fa85f64-5717-4562-b3fc-1c963f66ac33"
	registrationDT = "2025-04-14T12:33:30.752Z"
)

type LoginRequest struct {
	Role string `json:"role"`
}

type PvzRequest struct {
	ID               string `json:"id"`
	RegistrationDate string `json:"registrationDate"`
	City             string `json:"city"`
}

type ReceptionRequest struct {
	PvzID string `json:"pvzId"`
}

type ProductRequest struct {
	Type  string `json:"type"`
	PvzID string `json:"pvzId"`
}

func loginAndGetToken(t *testing.T, role string) string {
	t.Helper()

	payload := LoginRequest{Role: role}
	token, err := sendJSONRequest(loginURL, http.MethodPost, "", payload, http.StatusOK)
	if err != nil {
		t.Fatalf("Ошибка логина для %s: %v", role, err)
	}

	var result string
	if err := json.Unmarshal(token, &result); err != nil {
		t.Fatalf("Ошибка декодирования токена для %s: %v", role, err)
	}

	return result
}

func sendJSONRequest(url, method, token string, data interface{}, expectedStatus int) ([]byte, error) {
	var body *bytes.Reader
	if data != nil {
		bytesData, err := json.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("ошибка маршалинга: %v", err)
		}
		body = bytes.NewReader(bytesData)
	} else {
		body = bytes.NewReader([]byte{})
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания запроса: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != expectedStatus {
		return nil, fmt.Errorf("ожидался статус %d, получен %d", expectedStatus, resp.StatusCode)
	}

	respBody := new(bytes.Buffer)
	if _, err := respBody.ReadFrom(resp.Body); err != nil {
		return nil, fmt.Errorf("ошибка чтения ответа: %v", err)
	}

	return respBody.Bytes(), nil
}

func TestReceptionFlow(t *testing.T) {
	employeeToken := loginAndGetToken(t, "employee")
	moderatorToken := loginAndGetToken(t, "moderator")

	// 1. Отправляем запрос создания ПВЗ
	pvzData := PvzRequest{
		ID:               testPvzID,
		RegistrationDate: registrationDT,
		City:             "Moscow",
	}
	_, err := sendJSONRequest(apiPvzURL, http.MethodPost, moderatorToken, pvzData, http.StatusCreated)
	if err != nil {
		t.Fatalf("Ошибка создания ПВЗ: %v", err)
	}

	// 2. Отправляем запрос на создание приёмки
	receptionData := ReceptionRequest{
		PvzID: testPvzID,
	}
	t.Logf("Отправляемые данные для приёмки: %+v", receptionData)
	_, err = sendJSONRequest(apiReceptions, http.MethodPost, employeeToken, receptionData, http.StatusOK)
	if err != nil {
		t.Fatalf("Ошибка создания приёмки: %v", err)
	}

	// 3. Циклическое добавление продуктов (50 итераций)
	productData := ProductRequest{
		Type:  "electronics",
		PvzID: testPvzID,
	}
	for i := 1; i <= 50; i++ {
		_, err = sendJSONRequest(apiProducts, http.MethodPost, employeeToken, productData, http.StatusCreated)
		if err != nil {
			t.Fatalf("Ошибка создания продукта №%d: %v", i, err)
		}
	}

	// 4. Закрываем последнюю приёмку
	closeURL := fmt.Sprintf("%s%s/%s/close_last_reception", baseURL, pvzEndpoint, testPvzID)
	_, err = sendJSONRequest(closeURL, http.MethodPost, employeeToken, nil, http.StatusOK)
	if err != nil {
		t.Fatalf("Ошибка при закрытии последней приёмки: %v", err)
	}
}
