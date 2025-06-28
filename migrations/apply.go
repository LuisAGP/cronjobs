package migrations

import (
	"github.com/LuisAGP/cronjobs/app/models"
	"github.com/LuisAGP/cronjobs/app/services"
)

func ApplyMigrations() {
	db := services.DB()

	// Migraciones
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.User{},
	)
}
