package services

import (
	"ProjectFirst/models"
)

type ProductService interface {
	GetProductAll() ([]models.Product, error)
}
