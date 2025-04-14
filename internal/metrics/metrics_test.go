package metrics_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Wammero/PVZ-service/internal/metrics"
)

func TestMetricsMiddleware(t *testing.T) {
	handler := metrics.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Millisecond)
		w.WriteHeader(http.StatusTeapot)
		w.Write([]byte("I'm a teapot"))
	}))

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusTeapot {
		t.Errorf("Ожидался статус %d, получен %d", http.StatusTeapot, resp.StatusCode)
	}

	if !strings.Contains(string(body), "I'm a teapot") {
		t.Errorf("Ожидалось тело 'I'm a teapot', получено: %s", string(body))
	}
}

func TestWriteHeader(t *testing.T) {
	w := httptest.NewRecorder()
	rr := &mockResponseRecorder{ResponseWriter: w}

	rr.WriteHeader(http.StatusBadGateway)

	if rr.statusCode != http.StatusBadGateway {
		t.Errorf("Ожидался код %d, получен %d", http.StatusBadGateway, rr.statusCode)
	}
}

type mockResponseRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rr *mockResponseRecorder) WriteHeader(code int) {
	rr.statusCode = code
	rr.ResponseWriter.WriteHeader(code)
}

func TestMetricsHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/metrics", nil)
	w := httptest.NewRecorder()

	metrics.Handler().ServeHTTP(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Ожидался 200 OK, получен %d", resp.StatusCode)
	}

	if !strings.Contains(string(body), "http_requests_total") {
		t.Errorf("Метрика http_requests_total не найдена в выводе")
	}
}
