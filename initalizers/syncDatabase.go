package initalizers

import (
	"github.com/itismrx/jwt-auth/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
