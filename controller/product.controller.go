package controller

import (
	"ProjectFirst/services"
	"net/http"

	"github.com/labstack/echo"
)

var productService services.ProductService = services.ProductServiceImp{}

func SetProduct(c *echo.Group) {
	c.GET("/product", GetProductAll)
	c.GET("/product/:id", GetProductByID)
}

func GetProductAll(c echo.Context) error {

	data, err := productService.GetProductAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

func GetProductByID(c echo.Context) error {
	id := c.Param("id")

	result, err := productService.GetProductByID(id)
	if err != nil {
		return resErr(c, err)
	}

	return res(c, result)
}
