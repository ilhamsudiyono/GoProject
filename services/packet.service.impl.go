package services

import (
	"ProjectFirst/dao"
	"ProjectFirst/models"
)

var pcktDao dao.PacketDao = dao.PacketDaoImpl{}

type PacketServiceImp struct{}

func (PacketServiceImp) GetPacketByProductDtlId(id string) ([]models.Packet, error) {

	return pcktDao.GetPacketByProductDtlId(id)
}
