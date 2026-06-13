package departments

import "time"

// modelo base de departamento
type Department struct {
	ID             uint   `gorm:"primaryKey"`
	BusinessUnitID uint   `gorm:"not null;uniqueIndex:idx_department_business_slug"`
	Name           string `gorm:"type:varchar(150);not null"`
	Slug           string `gorm:"type:varchar(100);not null;uniqueIndex:idx_department_business_slug"`
	IsActive       bool   `gorm:"default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
