package model

type ProductType string

const (
	TypeElectronics ProductType = "электроника"
	TypeClothing    ProductType = "одежда"
)

var validProductTypes = map[ProductType]struct{}{
	TypeElectronics: {},
	TypeClothing:    {},
}

func IsValidProductType(pt ProductType) bool {
	_, ok := validProductTypes[pt]
	return ok
}

type UserRole string

const (
	RoleModerator UserRole = "moderator"
	RoleEmployee  UserRole = "employee"
)

var validUserRoles = map[UserRole]struct{}{
	RoleModerator: {},
	RoleEmployee:  {},
}

func IsValidUserRole(role UserRole) bool {
	_, ok := validUserRoles[role]
	return ok
}
