package services

import (
	"ProjectFirst/models"

	"gorm.io/gorm"
)

type TertanggungService interface {
	CreateTertanggung(tr *models.Tertanggung) (*models.Tertanggung, error)
	CreateTertanggung2nd(tr *models.Tertanggung, tx *gorm.DB) error
}
