package main

import (
	"github.com/joho/godotenv"
	"myponyasia.com/hub-api-seeder/models"
)

func main() {

	godotenv.Load()

	ConnectDB()

	var roles = []models.Role{
		{
			Name:        "ADMIN",
			Description: "Administrator",
			Level:       1,
		},
		{
			Name:        "MODERATOR",
			Description: "Moderator",
			Level:       2,
		},

		{
			Name:        "CREATOR",
			Description: "Creator",
			Level:       10,
		},
		{
			Name:        "PUBLISHER",
			Description: "Publisher",
			Level:       10,
		},

		{
			Name:        "USER",
			Description: "User",
			Level:       99,
		},
	}
	DB.Create(&roles)
}
