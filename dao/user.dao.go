package dao

import (
	model "ProjectFirst/models"
)

type UserDao interface {
	CreateUser(data *model.Users) (*model.Users, error)
	GetUserById(id string) (model.Users, error)
	GetUserByUsername(username string) (model.Users, error)
}
