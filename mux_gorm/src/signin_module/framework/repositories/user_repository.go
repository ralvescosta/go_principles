package signin_repositories

import (
	"fmt"
	migrations "gomux_gorm/src/core/database/table_models"
	entities "gomux_gorm/src/signin_module/bussiness/entities"

	"github.com/jinzhu/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type IUserRepository interface {
	Create(registerUser *entities.RegisterUsersEntity)
	FindOne(id int) *migrations.Users
	FindAll() *[]migrations.Users
	Update(id int) *migrations.Users
	Delete(id int) *migrations.Users
}

func (r *userRepository) Create(registerUser *entities.RegisterUsersEntity) {

	result := r.db.Create(registerUser)

	fmt.Println(result)
}

func (r *userRepository) FindOne(id int) *migrations.Users {
	user := migrations.Users{}

	result := r.db.First(&user, id)
	fmt.Println(result)

	return &user
}

func (r *userRepository) FindAll() *[]migrations.Users {

	user := []migrations.Users{migrations.Users{}}

	result := r.db.Find(&user)
	fmt.Println(result)

	return &user
}

func (r *userRepository) Update(id int) *migrations.Users {
	user := migrations.Users{}

	return &user
}

func (r *userRepository) Delete(id int) *migrations.Users {
	user := migrations.Users{}

	return &user
}

func UserRepositoryConstructor(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}
