package controller

import (
	"ProjectFirst/config"
	"ProjectFirst/models"
	"ProjectFirst/services"

	"github.com/labstack/echo"
)

var tertanggungPmSvc services.TertanggungPmService = services.TertanggungPmServiceImp{}

func SetTertanggungPm(c *echo.Group) {
	c.POST("/tertanggung-pm", PostTertanggungPm)
	// c.GET("/business-partner/:id", GetBusinessPartnerById)
	// c.GET("/business-partner/userid/:userid", GetBusinessPartnerByUserId)

}

func PostTertanggungPm(c echo.Context) (e error) {
	defer config.CatchError(&e)
	trtgPm := new(models.TertanggungPmPojo)

	if err := c.Bind(trtgPm); err != nil {
		return resErr(c, err)
	}
	var err = tertanggungPmSvc.CreateTertanggungPm(&trtgPm.Tertanggung, &trtgPm.PenerimaManfaat)

	if err == nil {
		return res(c, trtgPm)
	}
	return resErr(c, err)

}
