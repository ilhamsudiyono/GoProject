package controller

import (
	"ProjectFirst/config"
	"ProjectFirst/services"
	"log"

	"github.com/labstack/echo"
)

var packetDtlSvc services.PacketDtlService = services.PacketDtlServiceImp{}

func SetPacketDtl(c *echo.Group) {
	c.GET("/packet-detail/packetid/:packetId", GetPacketDtlByPacketId)

}

func GetPacketDtlByPacketId(c echo.Context) (e error) {
	defer config.CatchError(&e)

	id := c.Param("packetId")

	log.Println("=======CTRL PACKET==========")
	var result, err = packetDtlSvc.GetPacketDtlByPacketId(id)

	if err == nil {
		return res(c, result)

	}

	return resErr(c, err)
}
