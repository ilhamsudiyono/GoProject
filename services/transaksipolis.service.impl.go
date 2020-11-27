package services

import (
	"ProjectFirst/config"
	"ProjectFirst/dao"
	"ProjectFirst/models"
	"time"

	"gorm.io/gorm"
)

var trxPolisDao dao.TransaksiPolisDao = dao.TransaksiPolisDaoImpl{}

type TransaksiPolisServiceImp struct{}

func (TransaksiPolisServiceImp) CreateTransaksiPolis(txpl *models.TransaksiPolis) (*models.TransaksiPolis, error) {

	var createDt = models.Timestamp(time.Now())
	txpl.CreatedDate = &createDt

	return trxPolisDao.CreateTransaksiPolis(txpl)
}

func (TransaksiPolisServiceImp) CreateTransaksiPolis2nd(data *models.TransaksiPolis, tx *gorm.DB) (e error) {
	defer config.CatchError(&e)
	return trxPolisDao.CreateTransaksiPolis2nd(data, tx)
}
