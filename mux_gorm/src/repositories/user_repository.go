package repositories

import (
	"fmt"
	migrations "gomux_gorm/src/framework/database/migrations"

	"github.com/jinzhu/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type IUserRepository interface {
	FindOne(id int) migrations.Users
}

func (r *userRepository) FindOne(id int) migrations.Users {
	user := migrations.Users{}
	result := r.db.First(&user, id)

	fmt.Println(result)

	return user
}

func UserRepositoryConstructor(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}
