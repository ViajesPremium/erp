package businessunit

import "time"

type BusinessUnitSlug string

const (
	BusinessUnitMomentto BusinessUnitSlug = "momentto-garden"
	BusinessUnitViajes   BusinessUnitSlug = "viajes-premium"
)

// modelo base del negocio
type BusinessUnit struct {
	ID       uint             `gorm:"primaryKey"`
	Name     string           `gorm:"type:varchar(150);not null"`
	Slug     BusinessUnitSlug `gorm:"type:varchar(100);uniqueIndex;not null"`
	IsActive bool             `gorm:"default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
