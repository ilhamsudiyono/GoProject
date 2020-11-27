package dao

import (
	"ProjectFirst/config"
	"ProjectFirst/models"

	"gorm.io/gorm"
)

// declaration struct userdaoimpl
type TransaksiPolisDaoImpl struct{}

// create function for createUser
func (TransaksiPolisDaoImpl) CreateTransaksiPolis(trxpol *models.TransaksiPolis) (txpl *models.TransaksiPolis, e error) {
	defer config.CatchError(&e)
	// initiate result and inject to DB
	result := g.Create(trxpol)

	// return true
	if result.Error == nil {
		return trxpol, nil
	}

	// return error
	return nil, result.Error

}

func (TransaksiPolisDaoImpl) CreateTransaksiPolis2nd(data *models.TransaksiPolis, tx *gorm.DB) (e error) {
	defer config.CatchError(&e)
	if err := tx.Create(data).Error; err != nil {
		return err
	}
	return nil
}
