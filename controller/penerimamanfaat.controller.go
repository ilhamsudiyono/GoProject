package controller

import (
	"ProjectFirst/models"
	"ProjectFirst/responses"
	"ProjectFirst/services"

	"github.com/labstack/echo"
)

var pmft services.PenerimaManfaatService = services.PenerimaManfaatServiceImp{}

func SetPenerimaManfaat(c *echo.Group) {
	c.POST("/penerima-manfaat", PostPenerimaManfaat)
	// c.GET("/business-partner/:id", GetBusinessPartnerById)
	// c.GET("/business-partner/userid/:userid", GetBusinessPartnerByUserId)

}

func PostPenerimaManfaat(c echo.Context) error {
	pm := new(models.PenerimaManfaat)

	if err := c.Bind(pm); err != nil {
		return resErr(c, err)
	}

	result, err := pmft.CreatePenerimaManfaat(pm)

	if err != nil {
		return resErr(c, err)
	}

	rs := responses.NewPenerimaManfaat(result)
	return res(c, rs)

}
