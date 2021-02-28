package repositories

import (
	"gorm.io/gorm"
	"log"
	"user-go-test/model"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func CreateUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

func (u UserRepositoryImpl) GetAllUser() (*[]model.User, error) {
	var users []model.User
	err := u.db.Set("gorm:auto_preload",true).Find(&users).Error
	if err != nil {
		log.Fatal(err)
	}
	return &users, err
}

func (u UserRepositoryImpl) InsertNewUser(user *model.User) (*model.User, error) {
	err := u.db.Create(&user).Error
	if err != nil {
		log.Fatal(err)
	}
	return user, err
}

func (u UserRepositoryImpl) UpdateUser(id int, user *model.User) (*model.User, error) {
	err := u.db.Model(&user).Where("id=?",id).Updates(user).Error
	if err != nil {
		log.Fatal(err)
	}
	return user, err
}

func (u UserRepositoryImpl) GetUser(id int) (*model.User, error) {
	var user model.User
	err := u.db.Set("gorm:auto_preload",true).First(&user,id).Error
	if err != nil {
		log.Fatal(err)
	}
	return &user, err
}
