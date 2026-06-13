package roles

import "time"

type Roles string

const (
	RoleAdmin   Roles = "admin"
	RoleSeller  Roles = "seller"
	RoleSystems Roles = "systems"
)

// catalogo de roles por unidad de negocio
type Role struct {
	ID             uint  `gorm:"primaryKey"`
	BusinessUnitID uint  `gorm:"not null;uniqueIndex:idx_role_business_name"`
	Name           Roles `gorm:"type:varchar(100);not null;uniqueIndex:idx_role_business_name"`
	IsActive       bool  `gorm:"default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
