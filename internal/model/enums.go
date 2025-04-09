package model

type ProductType string

const (
	TypeElectronics ProductType = "электроника"
	TypeClothing    ProductType = "одежда"
)

func IsValidProductType(pt ProductType) bool {
	switch pt {
	case TypeElectronics, TypeClothing:
		return true
	}
	return false
}

type UserRole string

const (
	RoleClient    UserRole = "client"
	RoleModerator UserRole = "moderator"
	RoleEmployee  UserRole = "employee"
)

func IsValidUserRole(role UserRole) bool {
	switch role {
	case RoleClient, RoleModerator, RoleEmployee:
		return true
	}
	return false
}
