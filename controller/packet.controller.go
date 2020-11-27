package controller

import (
	"ProjectFirst/config"
	"ProjectFirst/services"
	"log"

	"github.com/labstack/echo"
)

var packetSvc services.PacketService = services.PacketServiceImp{}

func SetPacket(c *echo.Group) {
	c.GET("/packet/productid/:productDtlId", GetPacketByProductDtlId)

}

func GetPacketByProductDtlId(c echo.Context) (e error) {
	defer config.CatchError(&e)

	id := c.Param("productDtlId")

	log.Println("=======CTRL PACKET==========")
	var result, err = packetSvc.GetPacketByProductDtlId(id)

	if err == nil {
		return res(c, result)

	}

	return resErr(c, err)
}
