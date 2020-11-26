package services

import (
	"ProjectFirst/config"
	"ProjectFirst/dao"
	"ProjectFirst/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

var penerimamanfaatDao dao.PenerimaManfaatDao = dao.PenerimaManfaatDaoImpl{}

type PenerimaManfaatServiceImp struct{}

func (PenerimaManfaatServiceImp) CreatePenerimaManfaat(pm *models.PenerimaManfaat) (*models.PenerimaManfaat, error) {

	_, err := bpDao.GetBpById(pm.BusinessPartnerId)

	if err != nil {
		return nil, errors.New("BusinessPartner id not exist")
	}

	var createDt = models.Timestamp(time.Now())
	pm.CreatedDate = &createDt

	return penerimamanfaatDao.CreatePenerimaManfaat(pm)
}

func (PenerimaManfaatServiceImp) CreatePenerimaManfaat2nd(data *models.PenerimaManfaat, tx *gorm.DB) (e error) {
	defer config.CatchError(&e)
	return penerimamanfaatDao.CreatePenerimaManfaat2nd(data, tx)
}
