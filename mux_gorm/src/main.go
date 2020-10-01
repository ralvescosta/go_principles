package main

import (
	database "gomux_gorm/src/framework/database"
	"gomux_gorm/src/repositories"
)

func main() {
	db := database.ConnectToDatabase()

	userRepo := repositories.UserRepositoryConstructor(db)

	userRepo.FindOne(1)

}
