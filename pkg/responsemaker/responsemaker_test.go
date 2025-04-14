package responsemaker_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Wammero/PVZ-service/internal/model"
	"github.com/Wammero/PVZ-service/pkg/responsemaker"
)

func TestWriteJSONResponse(t *testing.T) {
	rr := httptest.NewRecorder()
	data := model.Error{Message: "что-то пошло не так"}

	responsemaker.WriteJSONResponse(rr, data, http.StatusBadRequest)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Ожидался статус %d, получен %d", http.StatusBadRequest, rr.Code)
	}

	if ct := rr.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("Ожидался Content-Type application/json, получен %s", ct)
	}

	var response model.Error
	err := json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Ошибка при декодировании JSON: %v", err)
	}

	if response.Message != data.Message {
		t.Errorf("Ожидалось сообщение %q, получено %q", data.Message, response.Message)
	}
}

func TestWriteJSONError(t *testing.T) {
	rr := httptest.NewRecorder()

	responsemaker.WriteJSONError(rr, "ошибка логина", http.StatusUnauthorized)

	if rr.Code != http.StatusUnauthorized {
		t.Errorf("Ожидался статус %d, получен %d", http.StatusUnauthorized, rr.Code)
	}

	var response model.Error
	err := json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Ошибка при декодировании JSON: %v", err)
	}

	if !strings.Contains(response.Message, "ошибка логина") {
		t.Errorf("Ожидалось сообщение, содержащее %q, получено %q", "ошибка логина", response.Message)
	}
}
