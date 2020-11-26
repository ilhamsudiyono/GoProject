package dao

import (
	"ProjectFirst/config"
	"ProjectFirst/models"

	"gorm.io/gorm"
)

// declaration struct userdaoimpl
type TertanggungDaoImpl struct{}

// create function for createUser
func (TertanggungDaoImpl) CreateTertanggung(tertanggung *models.Tertanggung) (trtg *models.Tertanggung, e error) {
	defer config.CatchError(&e)
	// initiate result and inject to DB
	result := g.Create(tertanggung)

	// return true
	if result.Error == nil {
		return tertanggung, nil
	}

	// return error
	return nil, result.Error

}

func (TertanggungDaoImpl) CreateTertanggung2nd(data *models.Tertanggung, tx *gorm.DB) (e error) {
	defer config.CatchError(&e)
	if err := tx.Create(data).Error; err != nil {
		return err
	}
	return nil
}
