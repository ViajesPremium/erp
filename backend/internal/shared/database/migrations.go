package database

// migrations.go es el archivo que debe encargarse de aplicar las migraciones de la base de datos.
import (
	businessunit "backend/internal/platform/business_unit"
	"backend/internal/platform/departments"
	"backend/internal/platform/roles"
	"backend/internal/platform/users"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&users.User{},
		&businessunit.BusinessUnit{},
		&businessunit.UserBusinessUnit{},
		&departments.Department{},
		&departments.UserDepartment{},
		&roles.Role{},
	)
}
