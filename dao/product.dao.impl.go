package dao

import (
	"ProjectFirst/models"
)

type ProductDaoImpl struct{}

func (ProductDaoImpl) GetProductAll() ([]models.Product, error) {
	var (
		data []models.Product
	)
	result := g.Order("id").Find(&data)

	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

func (ProductDaoImpl) GetProductByID(id string) (models.Product, error) {
	data := models.Product{}
	result := g.Where("id", id).First(&data)
	if result.Error != nil {
		return data, result.Error
	}

	return data, nil
}
