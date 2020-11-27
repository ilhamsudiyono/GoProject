package services

import (
	"ProjectFirst/models"

	"gorm.io/gorm"
)

type TransaksiPolisService interface {
	CreateTransaksiPolis(tr *models.TransaksiPolis) (*models.TransaksiPolis, error)
	CreateTransaksiPolis2nd(tr *models.TransaksiPolis, tx *gorm.DB) error
}
