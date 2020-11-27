package services

import (
	"ProjectFirst/dao"
	"ProjectFirst/models"
)

var pcktDtlDao dao.PacketDtlDao = dao.PacketDtlDaoImpl{}

type PacketDtlServiceImp struct{}

func (PacketDtlServiceImp) GetPacketDtlByPacketId(id string) ([]models.PacketDtl, error) {

	return pcktDtlDao.GetPacketDtlByPacketId(id)
}
