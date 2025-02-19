package initializers

import (
	custommodels "jwtgen/customModels"
	"jwtgen/models"
)

func SyncDb() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&custommodels.Employee{})
}
