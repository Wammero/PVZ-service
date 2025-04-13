package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Wammero/PVZ-service/internal/model"
	"github.com/Wammero/PVZ-service/pkg/jwt"
	"github.com/Wammero/PVZ-service/pkg/responsemaker"
	jwtlib "github.com/golang-jwt/jwt/v4"
)

func JWTValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			responsemaker.WriteJSONError(w, "Неавторизован.", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			responsemaker.WriteJSONError(w, "Неверный формат токена.", http.StatusUnauthorized)
			return
		}

		tokenStr := parts[1]
		claims := &model.Claims{}

		token, err := jwtlib.ParseWithClaims(tokenStr, claims, func(token *jwtlib.Token) (interface{}, error) {
			return jwt.GetSecret(), nil
		})
		if err != nil || !token.Valid {
			responsemaker.WriteJSONError(w, "Неверный токен.", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), model.UserIDContextKey, claims.UserID)
		ctx = context.WithValue(ctx, model.RoleContextKey, claims.Role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
