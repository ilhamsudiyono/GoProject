package controller

import (
	"ProjectFirst/config"
	"ProjectFirst/models"
	"ProjectFirst/services"

	"github.com/labstack/echo"
)

var trxPolisSvc services.TransaksiPolisService = services.TransaksiPolisServiceImp{}

func SetTrxPolis(c *echo.Group) {
	c.POST("/transaksi-polis", PostTransaksiPolis)
	// c.GET("/business-partner/:id", GetBusinessPartnerById)
	// c.GET("/business-partner/userid/:userid", GetBusinessPartnerByUserId)

}

func PostTransaksiPolis(c echo.Context) (e error) {
	defer config.CatchError(&e)
	trxPolis := new(models.TransaksiPolis)

	if err := c.Bind(trxPolis); err != nil {
		return resErr(c, err)
	}

	result, err := trxPolisSvc.CreateTransaksiPolis(trxPolis)

	if err != nil {
		return resErr(c, err)
	}

	return res(c, result)

}
