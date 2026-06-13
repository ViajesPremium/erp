package users

import "time"

// int permite negativos y positivos
// uint solo permite positivos y cero

// modelo base del usuario
type User struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"type:varchar(150);not null"`
	Email        string `gorm:"type:varchar(150);uniqueIndex;not null"`
	PasswordHash string `gorm:"type:text;not null"`
	IsActive     bool   `gorm:"default:true"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
