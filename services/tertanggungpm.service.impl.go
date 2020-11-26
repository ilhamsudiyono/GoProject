package services

import (
	"ProjectFirst/config"
	"ProjectFirst/models"

	"gorm.io/gorm"
)

// var tertggDao dao.TertanggungDao = dao.TertanggungDaoImpl{}

type TertanggungPmServiceImp struct{}

func (TertanggungPmServiceImp) CreateTertanggungPm(trtgArr *[]models.Tertanggung, pmArr *[]models.PenerimaManfaat) (e error) {
	defer config.CatchError(&e)

	return g.Transaction(func(tx *gorm.DB) error {
		for _, v := range *trtgArr {

			if err := tertggDao.CreateTertanggung2nd(&v, tx); err != nil {
				return err
			}
		}
		for _, w := range *pmArr {

			if err := penerimamanfaatDao.CreatePenerimaManfaat2nd(&w, tx); err != nil {
				return err
			}
		}

		return nil

	})
}
