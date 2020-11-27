package services

import (
	"ProjectFirst/models"
)

type PacketDtlService interface {
	GetPacketDtlByPacketId(id string) ([]models.PacketDtl, error)
}
