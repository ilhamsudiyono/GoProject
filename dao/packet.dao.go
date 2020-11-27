package dao

import (
	"ProjectFirst/models"
)

// declaration interface productdao
type PacketDao interface {
	GetPacketByProductDtlId(id string) ([]models.Packet, error)
}
