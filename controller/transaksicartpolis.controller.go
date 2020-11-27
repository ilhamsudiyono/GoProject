package controller

import (
	"ProjectFirst/config"
	"ProjectFirst/models"
	"ProjectFirst/services"

	"github.com/labstack/echo"
)

var trxCartPolisSvc services.TransaksiCartPolisService = services.TransaksiCartPolisServiceImp{}

func SetTrxCartPolis(c *echo.Group) {
	c.POST("/transaksi-cart-polis", PostTransaksiCartPolis)
	// c.GET("/business-partner/:id", GetBusinessPartnerById)
	// c.GET("/business-partner/userid/:userid", GetBusinessPartnerByUserId)

}

func PostTransaksiCartPolis(c echo.Context) (e error) {
	defer config.CatchError(&e)
	trxCrtPolis := new(models.TransaksiCartPolisPojo)

	if err := c.Bind(trxCrtPolis); err != nil {
		return resErr(c, err)
	}
	var err = trxCartPolisSvc.CreateTransaksiCartPolis(&trxCrtPolis.TransaksiCart, &trxCrtPolis.TransaksiPolis)

	if err == nil {
		return res(c, trxCrtPolis)
	}
	return resErr(c, err)

}
