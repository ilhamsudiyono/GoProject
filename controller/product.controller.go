package controller

import (
	"ProjectFirst/services"
	"net/http"

	"github.com/labstack/echo"
)

var productService services.ProductService = services.ProductServiceImp{}

func SetProduct(c *echo.Group) {
	c.GET("/product", GetProductAll)
}

func GetProductAll(c echo.Context) error {

	data, err := productService.GetProductAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}
