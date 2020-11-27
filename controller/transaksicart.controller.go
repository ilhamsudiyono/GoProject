package controller

import (
	"ProjectFirst/config"
	"ProjectFirst/models"
	"ProjectFirst/services"

	"github.com/labstack/echo"
)

var trxCartSvc services.TransaksiCartService = services.TransaksiCartServiceImp{}

func SetTrxCart(c *echo.Group) {
	c.POST("/transaksi-cart", PostTransaksiCart)
	c.GET("/transaksi-cart/:id", GetTransaksiCartById)
	// c.GET("/business-partner/userid/:userid", GetBusinessPartnerByUserId)

}

func PostTransaksiCart(c echo.Context) (e error) {
	defer config.CatchError(&e)
	trxCrt := new(models.TransaksiCart)

	if err := c.Bind(trxCrt); err != nil {
		return resErr(c, err)
	}

	result, err := trxCartSvc.CreateTransaksiCart(trxCrt)

	if err != nil {
		return resErr(c, err)
	}

	return res(c, result)

}

func GetTransaksiCartById(c echo.Context) error {
	id := c.Param("id")

	result, err := trxCartSvc.GetTransaksiCartByID(id)
	if err != nil {
		return resErr(c, err)
	}

	return res(c, result)
}
