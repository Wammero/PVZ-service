package model_test

import (
	"testing"

	"github.com/Wammero/PVZ-service/internal/model"
)

func TestIsValidProductType(t *testing.T) {
	tests := []struct {
		productType model.ProductType
		expected    bool
	}{
		{model.TypeElectronics, true},
		{model.TypeClothing, true},
		{"invalid", false},
	}

	for _, tt := range tests {
		t.Run(string(tt.productType), func(t *testing.T) {
			result := model.IsValidProductType(tt.productType)
			if result != tt.expected {
				t.Errorf("IsValidProductType(%v) = %v; want %v", tt.productType, result, tt.expected)
			}
		})
	}
}

func TestIsValidUserRole(t *testing.T) {
	tests := []struct {
		role     model.UserRole
		expected bool
	}{
		{model.RoleModerator, true},
		{model.RoleEmployee, true},
		{"invalid", false},
	}

	for _, tt := range tests {
		t.Run(string(tt.role), func(t *testing.T) {
			result := model.IsValidUserRole(tt.role)
			if result != tt.expected {
				t.Errorf("IsValidUserRole(%v) = %v; want %v", tt.role, result, tt.expected)
			}
		})
	}
}
