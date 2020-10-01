package main

import (
	database "gomux_gorm/src/core/database"
	repositories "gomux_gorm/src/signin_module/framework/repositories"
)

func main() {
	db := database.ConnectToDatabase()

	userRepo := repositories.UserRepositoryConstructor(db)

	userRepo.FindOne(1)

}
