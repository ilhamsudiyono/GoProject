package services

import (
	"ProjectFirst/models"
)

type PacketService interface {
	GetPacketByProductDtlId(id string) ([]models.Packet, error)
}
