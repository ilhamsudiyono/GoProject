package dao

import (
	"ProjectFirst/models"
)

// declaration struct userdaoimpl
type PacketDtlDaoImpl struct{}

// GetBpByUserId

func (PacketDtlDaoImpl) GetPacketDtlByPacketId(id string) ([]models.PacketDtl, error) {
	var data []models.PacketDtl

	result := g.Preload("Packet.ProductDtl.Product").Where("packet_id", id).Find(&data)

	if result.Error != nil {
		return data, result.Error
	}
	return data, nil

}
