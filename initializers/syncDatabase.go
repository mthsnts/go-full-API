package initializers

import "gihub.com/mthsnts/go-JWT-API/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
