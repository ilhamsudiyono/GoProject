package services

import (
	"ProjectFirst/dao"
	"ProjectFirst/models"
)

var productDao dao.ProductDao = dao.ProductDaoImpl{}

type ProductServiceImp struct{}

func (ProductServiceImp) GetProductAll() ([]models.Product, error) {
	return productDao.GetProductAll()
}

func (ProductServiceImp) GetProductByID(id string) (models.Product, error) {
	return productDao.GetProductByID(id)
}
