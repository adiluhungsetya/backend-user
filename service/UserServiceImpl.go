package service

import (
	"user-go-test/model"
	"user-go-test/repositories"
)

type UserServiceImpl struct {
	userRepository repositories.UserRepository
}

func CreateUserService (userRepository repositories.UserRepository) UserService {
	return &UserServiceImpl{userRepository}
}

func (u UserServiceImpl) GetAllUser() (*[]model.User, error) {
	return u.userRepository.GetAllUser()
}

func (u UserServiceImpl) InsertNewUser(user *model.User) (*model.User, error) {
	return u.userRepository.InsertNewUser(user)
}

func (u UserServiceImpl) UpdateUser(id int, user *model.User) (*model.User, error) {
	return u.userRepository.UpdateUser(id,user)
}

func (u UserServiceImpl) GetUser(id int) (*model.User, error) {
	return u.userRepository.GetUser(id)
}
