package middleware

import (
	"net/http"

	"github.com/Wammero/PVZ-service/pkg/jwt"
	"github.com/Wammero/PVZ-service/pkg/responsemaker"
)

func hasAccess(role string, allowedRoles ...string) bool {
	for _, r := range allowedRoles {
		if role == r {
			return true
		}
	}
	return false
}

func RequireRole(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			role, ok := jwt.GetUserRole(r.Context())
			if !ok || !hasAccess(role, allowedRoles...) {
				responsemaker.WriteJSONError(w, "доступ запрещён: недостаточно прав", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
