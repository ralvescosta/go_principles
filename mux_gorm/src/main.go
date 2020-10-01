package main

import (
	migrations "gomux_gorm/src/framework/database/migrations"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	connection, err := gorm.Open("postgres", "user=postgres password=12345 dbname=default sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	connection.AutoMigrate(&migrations.Users{}, &migrations.Permissions{}, &migrations.UserPermissions{}, &migrations.Sessions{})
	// connection.Create(&framework.Users{Email: "user@email.com", Name: "user", LastName: "last", Password: "123456"})

	database := connection.DB()
	err = database.Ping()
	if err != nil {
		log.Fatal(err)
	}

}
