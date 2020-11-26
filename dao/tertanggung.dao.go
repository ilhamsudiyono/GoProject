package dao

import (
	"ProjectFirst/models"

	"gorm.io/gorm"
)

// declaration interface userdao
type TertanggungDao interface {
	CreateTertanggung(data *models.Tertanggung) (*models.Tertanggung, error)

	CreateTertanggung2nd(data *models.Tertanggung, tx *gorm.DB) error
}
