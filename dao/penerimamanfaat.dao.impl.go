package dao

import (
	"ProjectFirst/config"
	"ProjectFirst/models"

	"gorm.io/gorm"
)

// declaration struct userdaoimpl
type PenerimaManfaatDaoImpl struct{}

// create function for createUser
func (PenerimaManfaatDaoImpl) CreatePenerimaManfaat(penerimaManfaat *models.PenerimaManfaat) (pm *models.PenerimaManfaat, e error) {
	defer config.CatchError(&e)
	// initiate result and inject to DB
	result := g.Create(penerimaManfaat)

	// return true
	if result.Error == nil {
		return penerimaManfaat, nil
	}

	// return error
	return nil, result.Error

}

func (PenerimaManfaatDaoImpl) CreatePenerimaManfaat2nd(data *models.PenerimaManfaat, tx *gorm.DB) (e error) {
	defer config.CatchError(&e)
	if err := tx.Create(data).Error; err != nil {
		return err
	}
	return nil
}
