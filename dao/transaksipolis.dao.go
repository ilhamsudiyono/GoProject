package dao

import (
	"ProjectFirst/models"

	"gorm.io/gorm"
)

// declaration interface userdao
type TransaksiPolisDao interface {
	CreateTransaksiPolis(data *models.TransaksiPolis) (*models.TransaksiPolis, error)
	CreateTransaksiPolis2nd(data *models.TransaksiPolis, tx *gorm.DB) error
}
