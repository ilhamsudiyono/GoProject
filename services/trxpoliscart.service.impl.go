package services

import (
	"ProjectFirst/config"
	"ProjectFirst/models"

	"gorm.io/gorm"
)

type TransaksiCartPolisServiceImp struct{}

func (TransaksiCartPolisServiceImp) CreateTransaksiCartPolis(cartgArr *[]models.TransaksiCart, polisArr *[]models.TransaksiPolis) (e error) {
	defer config.CatchError(&e)

	return g.Transaction(func(tx *gorm.DB) error {
		for _, v := range *cartgArr {

			if err := trxCartDao.CreateTransaksiCart2nd(&v, tx); err != nil {
				return err
			}
		}
		for _, w := range *polisArr {

			if err := trxPolisDao.CreateTransaksiPolis2nd(&w, tx); err != nil {
				return err
			}
		}

		return nil

	})
}
