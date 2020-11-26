package dao

import (
	"ProjectFirst/config"
	"ProjectFirst/models"
)

// declaration struct userdaoimpl
type BpDaoImpl struct{}

// create function for createUser
func (BpDaoImpl) CreateBp(businessPartner *models.BusinessPartner) (bp *models.BusinessPartner, e error) {
	defer config.CatchError(&e)
	// initiate result and inject to DB
	result := g.Create(businessPartner)

	// return true
	if result.Error == nil {
		return businessPartner, nil
	}

	// return error
	return nil, result.Error

}

func (BpDaoImpl) GetBpById(id string) (models.BusinessPartner, error) {
	// initiate data from model users
	var data models.BusinessPartner

	// initiate result and inject like query to db
	result := g.Preload("Users.Roles").Where("id", id).First(&data)

	// return error
	if result.Error != nil {
		return data, result.Error
	}

	//return true
	return data, nil
}

// GetBpByUserId

func (BpDaoImpl) GetBpByUserId(id string) (models.BusinessPartner, error) {
	var data models.BusinessPartner

	result := g.Preload("Users.Roles").Where("user_id", id).First(&data)

	// return error
	if result.Error != nil {
		return data, result.Error
	}

	//return true
	return data, nil

}
