package services

import (
	"ProjectFirst/models"

	"gorm.io/gorm"
)

type TransaksiCartService interface {
	CreateTransaksiCart(tr *models.TransaksiCart) (*models.TransaksiCart, error)
	CreateTransaksiCart2nd(tr *models.TransaksiCart, tx *gorm.DB) error
	GetTransaksiCartByID(id string) (models.TransaksiCart, error)
}
