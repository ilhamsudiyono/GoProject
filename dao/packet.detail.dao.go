package dao

import (
	"ProjectFirst/models"
)

// declaration interface productdao
type PacketDtlDao interface {
	GetPacketDtlByPacketId(id string) ([]models.PacketDtl, error)
}
