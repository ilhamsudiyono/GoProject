package dao

import (
	"ProjectFirst/config"
	model "ProjectFirst/models"
)

type UserDaoImpl struct{}

func (UserDaoImpl) CreateUser(user *model.Users) (u *model.Users, e error) {
	defer config.CatchError(&e)
	result := g.Create(user)

	if result.Error == nil {
		return user, nil
	}

	return nil, result.Error

}

func (UserDaoImpl) GetUserById(id string) (u model.Users, e error) {
	defer config.CatchError(&e)

	data := model.Users{}

	result := g.Where("id", id).First(&data)

	if result.Error != nil {
		return data, result.Error
	}

	return data, nil
}

func (UserDaoImpl) GetUserByUsername(username string) (u model.Users, e error) {
	defer config.CatchError(&e)

	var data model.Users
	result := g.Where("email = ?", username).Find(&data)

	if result.Error == nil {
		return data, nil
	}

	return data, result.Error
}
