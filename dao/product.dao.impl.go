package dao

import (
	"ProjectFirst/models"
)

type ProductDaoImpl struct{}

func (ProductDaoImpl) GetProductAll() ([]models.Product, error) {
	var (
		data []models.Product
	)
	result := g.Preload("ProductDtl").Order("id").Find(&data)

	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}
