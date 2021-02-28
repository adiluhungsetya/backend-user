package repositories

import "user-go-test/model"

type UserRepository interface {
	GetAllUser() (*[]model.User,error)
	InsertNewUser(user *model.User)(*model.User,error)
	UpdateUser(id int,user *model.User)(*model.User,error)
	GetUser(id int) (*model.User,error)
}