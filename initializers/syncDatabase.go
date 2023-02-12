package initializers

import "github.com/edutjie/go-crud/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{}, &models.Post{})
}
