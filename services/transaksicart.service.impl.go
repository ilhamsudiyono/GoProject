package services

import (
	"ProjectFirst/config"
	"ProjectFirst/dao"
	"ProjectFirst/models"
	"time"

	"gorm.io/gorm"
)

var trxCartDao dao.TransaksiCartDao = dao.TransaksiCartDaoImpl{}

type TransaksiCartServiceImp struct{}

func (TransaksiCartServiceImp) CreateTransaksiCart(txcrt *models.TransaksiCart) (*models.TransaksiCart, error) {

	var createDt = models.Timestamp(time.Now())
	txcrt.CreatedDate = &createDt

	return trxCartDao.CreateTransaksiCart(txcrt)
}

func (TransaksiCartServiceImp) CreateTransaksiCart2nd(data *models.TransaksiCart, tx *gorm.DB) (e error) {
	defer config.CatchError(&e)
	return trxCartDao.CreateTransaksiCart2nd(data, tx)
}

func (TransaksiCartServiceImp) GetTransaksiCartByID(id string) (models.TransaksiCart, error) {
	return trxCartDao.GetTransaksiCartByID(id)
}
