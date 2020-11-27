package dao

import (
	"ProjectFirst/config"
	"ProjectFirst/models"

	"gorm.io/gorm"
)

// declaration struct userdaoimpl
type TransaksiCartDaoImpl struct{}

// create function for createUser
func (TransaksiCartDaoImpl) CreateTransaksiCart(trxcart *models.TransaksiCart) (txct *models.TransaksiCart, e error) {
	defer config.CatchError(&e)
	// initiate result and inject to DB
	result := g.Create(trxcart)

	// return true
	if result.Error == nil {
		return trxcart, nil
	}

	// return error
	return nil, result.Error

}

func (TransaksiCartDaoImpl) CreateTransaksiCart2nd(data *models.TransaksiCart, tx *gorm.DB) (e error) {
	defer config.CatchError(&e)
	if err := tx.Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (TransaksiCartDaoImpl) GetTransaksiCartByID(id string) (models.TransaksiCart, error) {
	var data models.TransaksiCart

	result := g.Preload("BusinessPartner.Users.Roles").Where("id", id).First(&data)

	if result.Error != nil {
		return data, result.Error
	}
	return data, nil

}
