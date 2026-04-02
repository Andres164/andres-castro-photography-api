package models


const (
	RoleAdmin = "admin"
	RoleViewer = "viewer"
)

func IsValidRole(role string) bool  {
	switch role {
	case RoleAdmin, RoleViewer:
		return true
	default:
		return false
	}
}