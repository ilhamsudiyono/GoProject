package services

import (
	"ProjectFirst/models"
)

type BusinessPartnerService interface {
	CreateBusinessPartner(bp *models.BusinessPartner) (*models.BusinessPartner, error)
	GetBpById(id string) (models.BusinessPartner, error)
	GetBpByUserId(id string) (models.BusinessPartner, error)
}
