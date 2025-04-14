package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Wammero/PVZ-service/internal/model"
	"github.com/Wammero/PVZ-service/pkg/jwt"
	jwtlib "github.com/golang-jwt/jwt/v4"
)

func mockHandler(t *testing.T) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("OK"))
		if err != nil {
			t.Errorf("Ошибка при записи ответа: %v", err)
		}
	})
}

func generateTestToken(t *testing.T, userID int, role string, secret string) string {
	jwt.SetSecret(secret)
	claims := &model.Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwtlib.RegisteredClaims{
			ExpiresAt: jwtlib.NewNumericDate(time.Now().Add(time.Hour)),
			Issuer:    "PVZ-service",
		},
	}
	token := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		t.Fatalf("Ошибка генерации токена: %v", err)
	}
	return signedToken
}

func TestJWTValidator_Success(t *testing.T) {
	secret := "test-secret"
	token := generateTestToken(t, 42, "admin", secret)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()

	handler := JWTValidator(mockHandler(t))
	jwt.SetSecret(secret)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Ожидался статус 200, но получен %d", w.Code)
	}
}

func TestJWTValidator_MissingHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	handler := JWTValidator(mockHandler(t))
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Ожидался статус 401, но получен %d", w.Code)
	}
}

func TestJWTValidator_InvalidFormat(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "InvalidTokenFormat")
	w := httptest.NewRecorder()

	handler := JWTValidator(mockHandler(t))
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Ожидался статус 401, но получен %d", w.Code)
	}
}

func TestJWTValidator_InvalidToken(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer invalid.token.here")
	w := httptest.NewRecorder()

	handler := JWTValidator(mockHandler(t))
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Ожидался статус 401, но получен %d", w.Code)
	}
}

func TestHasAccess(t *testing.T) {
	tests := []struct {
		role         string
		allowedRoles []string
		expected     bool
	}{
		{"admin", []string{"admin", "user"}, true},
		{"user", []string{"admin"}, false},
		{"guest", []string{}, false},
	}

	for _, test := range tests {
		result := hasAccess(test.role, test.allowedRoles...)
		if result != test.expected {
			t.Errorf("hasAccess(%q, %v) = %v; ожидалось %v",
				test.role, test.allowedRoles, result, test.expected)
		}
	}
}
