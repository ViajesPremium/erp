package businessunit

import "time"

// modelo de relacion entre usuario y unidad de negocio
type UserBusinessUnit struct {
	ID             uint `gorm:"primaryKey"`
	UserID         uint `gorm:"not null;uniqueIndex:idx_user_business"`
	BusinessUnitID uint `gorm:"not null;uniqueIndex:idx_user_business"`

	IsActive bool `gorm:"default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
