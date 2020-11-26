package services

import (
	"ProjectFirst/config"
	"ProjectFirst/dao"
	"ProjectFirst/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

var tertggDao dao.TertanggungDao = dao.TertanggungDaoImpl{}

type TertanggungServiceImp struct{}

func (TertanggungServiceImp) CreateTertanggung(trtg *models.Tertanggung) (*models.Tertanggung, error) {

	_, err := userDao.GetUserById(trtg.UserId)

	if err != nil {
		return nil, errors.New("User id not exist")
	}

	var createDt = models.Timestamp(time.Now())
	trtg.CreatedDate = &createDt

	return tertggDao.CreateTertanggung(trtg)
}

func (TertanggungServiceImp) CreateTertanggung2nd(data *models.Tertanggung, tx *gorm.DB) (e error) {
	defer config.CatchError(&e)
	return tertggDao.CreateTertanggung2nd(data, tx)
}
