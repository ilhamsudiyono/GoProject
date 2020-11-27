package dao

import (
	"ProjectFirst/models"
)

// declaration struct userdaoimpl
type PacketDaoImpl struct{}

// GetBpByUserId

func (PacketDaoImpl) GetPacketByProductDtlId(id string) ([]models.Packet, error) {
	var data []models.Packet

	result := g.Preload("ProductDtl.Product").Where("product_dtl_id", id).Find(&data)

	// return error
	if result.Error != nil {
		return data, result.Error
	}

	//return true
	return data, nil

}
