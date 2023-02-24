package pivot

import "github.com/google/uuid"

// UserRoles represents the database model of user roles relationships
type UserRoles struct {
	UserID uuid.UUID `gorm:"primary_key" json:"user_id"`
	RoleID uint      `gorm:"primary_key" json:"role_id"`
}

// TableName sets the table name
func (UserRoles) TableName() string {
	return "user_roles"
}
