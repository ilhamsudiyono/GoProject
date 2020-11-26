package dao

import (
	"ProjectFirst/models"

	"gorm.io/gorm"
)

// declaration interface userdao
type PenerimaManfaatDao interface {
	CreatePenerimaManfaat(data *models.PenerimaManfaat) (*models.PenerimaManfaat, error)

	CreatePenerimaManfaat2nd(data *models.PenerimaManfaat, tx *gorm.DB) error
}
