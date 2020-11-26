package services

import (
	"ProjectFirst/models"

	"gorm.io/gorm"
)

type PenerimaManfaatService interface {
	CreatePenerimaManfaat(tr *models.PenerimaManfaat) (*models.PenerimaManfaat, error)
	CreatePenerimaManfaat2nd(tr *models.PenerimaManfaat, tx *gorm.DB) error
}
