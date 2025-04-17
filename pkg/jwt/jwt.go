package jwt

import (
	"context"
	"time"

	"github.com/Wammero/PVZ-service/internal/model"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret []byte

func SetSecret(secret string) {
	jwtSecret = []byte(secret)
}

func GetSecret() []byte {
	return jwtSecret
}

func GenerateJWT(userID string, role string) (string, error) {
	claims := &model.Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			Issuer:    "PVZ-service",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func GetUserID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(model.UserIDContextKey).(string)
	return id, ok
}

func GetUserRole(ctx context.Context) (string, bool) {
	role, ok := ctx.Value(model.RoleContextKey).(string)
	return role, ok
}
