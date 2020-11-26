package controller

import (
	"ProjectFirst/config"
	"ProjectFirst/models"
	"ProjectFirst/services"
	"log"

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
	log.Println(">>>>>", &trtgPm.Tertanggung)
	var err = tertanggungPmSvc.CreateTertanggungPm(&trtgPm.Tertanggung, &trtgPm.PenerimaManfaat)

	if err == nil {
		return res(c, trtgPm)
	}
	return resErr(c, err)

}
