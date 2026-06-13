package departments

import "time"

// relacion entre usuario, departamento y rol
type UserDepartment struct {
	ID uint `gorm:"primaryKey"`

	UserID       uint `gorm:"not null;uniqueIndex:idx_user_department"`
	DepartmentID uint `gorm:"not null;uniqueIndex:idx_user_department"`
	RoleID       uint `gorm:"not null;index"`

	IsActive bool `gorm:"default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
