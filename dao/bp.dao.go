package dao

import (
	"ProjectFirst/models"
)

// declaration interface userdao
type BpDao interface {
	CreateBp(data *models.BusinessPartner) (*models.BusinessPartner, error)
	GetBpById(id string) (models.BusinessPartner, error)
	GetBpByUserId(id string) (models.BusinessPartner, error)
}
