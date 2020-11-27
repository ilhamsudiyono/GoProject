package dao

import (
	"ProjectFirst/models"
)

// declaration interface productdao
type ProductDao interface {
	GetProductAll() ([]models.Product, error)
	GetProductByID(id string) (models.Product, error)
}
