package model

import "github.com/golang-jwt/jwt/v4"

type ContextKey string

const (
	UserIDContextKey ContextKey = "user_id"
	RoleContextKey   ContextKey = "role"
)

type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}
