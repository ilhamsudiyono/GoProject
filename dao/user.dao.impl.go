package dao

import (
	"ProjectFirst/config"
	model "ProjectFirst/models"
)

// declaration struct userdaoimpl
type UserDaoImpl struct{}

// create function for createUser
func (UserDaoImpl) CreateUser(user *model.Users) (u *model.Users, e error) {
	defer config.CatchError(&e)
	// initiate result and inject to DB
	result := g.Create(user)

	// return true
	if result.Error == nil {
		return user, nil
	}

	// return error
	return nil, result.Error

}

// create function for getUserById
func (UserDaoImpl) GetUserById(id string) (u model.Users, e error) {
	defer config.CatchError(&e)

	// initiate data from model users
	data := model.Users{}

	// initiate result and inject like query to db
	result := g.Where("id", id).First(&data)

	// return error
	if result.Error != nil {
		return data, result.Error
	}

	//return true
	return data, nil
}

// create function for getUserByUsername
func (UserDaoImpl) GetUserByUsername(username string) (u model.Users, e error) {
	defer config.CatchError(&e)

	// initiate data from model users
	var data model.Users
	// initiate result and inject like query to db
	result := g.Where("email = ?", username).Find(&data)

	// return true
	if result.Error == nil {
		return data, nil
	}

	// return false
	return data, result.Error
}
