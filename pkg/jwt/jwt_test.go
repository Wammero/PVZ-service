package jwt

import (
	"context"
	"testing"
	"time"

	"github.com/Wammero/PVZ-service/internal/model"
	"github.com/golang-jwt/jwt/v4"
)

func TestSetAndGetSecret(t *testing.T) {
	secret := "mysecret"
	SetSecret(secret)

	if string(GetSecret()) != secret {
		t.Errorf("Expected secret %s, got %s", secret, string(GetSecret()))
	}
}

func TestGenerateJWT(t *testing.T) {
	SetSecret("mysecret")

	userID := 123
	role := "admin"

	tokenStr, err := GenerateJWT(userID, role)
	if err != nil {
		t.Fatalf("GenerateJWT returned error: %v", err)
	}

	if tokenStr == "" {
		t.Fatal("GenerateJWT returned empty token")
	}

	token, err := jwt.ParseWithClaims(tokenStr, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetSecret(), nil
	})
	if err != nil {
		t.Fatalf("Failed to parse token: %v", err)
	}

	if !token.Valid {
		t.Fatal("Parsed token is not valid")
	}

	claims, ok := token.Claims.(*model.Claims)
	if !ok {
		t.Fatal("Claims type assertion failed")
	}

	if claims.UserID != userID {
		t.Errorf("Expected UserID %d, got %d", userID, claims.UserID)
	}

	if claims.Role != role {
		t.Errorf("Expected Role %s, got %s", role, claims.Role)
	}

	if claims.Issuer != "PVZ-service" {
		t.Errorf("Expected Issuer PVZ-service, got %s", claims.Issuer)
	}

	if time.Until(claims.ExpiresAt.Time) > time.Hour || time.Until(claims.ExpiresAt.Time) < time.Minute*59 {
		t.Errorf("Unexpected expiration time: %v", claims.ExpiresAt.Time)
	}
}

func TestGetUserID(t *testing.T) {
	ctx := context.WithValue(context.Background(), model.UserIDContextKey, 99)
	id, ok := GetUserID(ctx)

	if !ok {
		t.Error("Expected to find user ID in context, but didn't")
	}
	if id != 99 {
		t.Errorf("Expected ID 99, got %d", id)
	}
}

func TestGetUserRole(t *testing.T) {
	ctx := context.WithValue(context.Background(), model.RoleContextKey, "manager")
	role, ok := GetUserRole(ctx)

	if !ok {
		t.Error("Expected to find role in context, but didn't")
	}
	if role != "manager" {
		t.Errorf("Expected role 'manager', got '%s'", role)
	}
}
