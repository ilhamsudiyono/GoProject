package controller

import (
	"ProjectFirst/models"
	"ProjectFirst/responses"
	"ProjectFirst/services"

	"github.com/labstack/echo"
)

var tertgg services.TertanggungService = services.TertanggungServiceImp{}

func SetTertanggung(c *echo.Group) {
	c.POST("/tertanggung", PostTertanggung)
	// c.GET("/business-partner/:id", GetBusinessPartnerById)
	// c.GET("/business-partner/userid/:userid", GetBusinessPartnerByUserId)

}

func PostTertanggung(c echo.Context) error {
	trtgg := new(models.Tertanggung)

	if err := c.Bind(trtgg); err != nil {
		return resErr(c, err)
	}

	result, err := tertgg.CreateTertanggung(trtgg)

	if err != nil {
		return resErr(c, err)
	}

	rs := responses.NewTertanggung(result)
	return res(c, rs)

}
