package initializers

import "v1/src/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
