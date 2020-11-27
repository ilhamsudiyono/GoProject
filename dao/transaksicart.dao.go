package dao

import (
	"ProjectFirst/models"

	"gorm.io/gorm"
)

// declaration interface userdao
type TransaksiCartDao interface {
	CreateTransaksiCart(data *models.TransaksiCart) (*models.TransaksiCart, error)
	CreateTransaksiCart2nd(data *models.TransaksiCart, tx *gorm.DB) error
	GetTransaksiCartByID(id string) (models.TransaksiCart, error)
}
