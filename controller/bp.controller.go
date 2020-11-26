package controller

import (
	"ProjectFirst/config"
	"ProjectFirst/requests"
	"ProjectFirst/responses"
	"ProjectFirst/services"

	"github.com/labstack/echo"
)

var businessPartnerSvc services.BusinessPartnerService = services.BpServiceImp{}

func SetBp(c *echo.Group) {
	c.POST("/business-partner", PostBusinessPartner)
	c.GET("/business-partner/:id", GetBusinessPartnerById)
	c.GET("/business-partner/userid/:userid", GetBusinessPartnerByUserId)

}

func PostBusinessPartner(c echo.Context) error {
	req := &requests.BusinessPartnerRequest{}
	req = req.Convert(c)

	result, err := businessPartnerSvc.CreateBusinessPartner(req.Model())
	if err != nil {
		return resErr(c, err)
	}

	rs := responses.NewBp(result)
	return res(c, rs)
}

func GetBusinessPartnerById(c echo.Context) (e error) {
	defer config.CatchError(&e)

	id := c.Param("id")
	var result, err = businessPartnerSvc.GetBpById(id)

	if err == nil {
		return res(c, result)

	}

	return resErr(c, err)

}

func GetBusinessPartnerByUserId(c echo.Context) (e error) {
	defer config.CatchError(&e)

	id := c.Param("userid")

	var result, err = businessPartnerSvc.GetBpByUserId(id)

	if err == nil {
		return res(c, result)

	}

	return resErr(c, err)
}
