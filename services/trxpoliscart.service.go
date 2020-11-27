package services

import "ProjectFirst/models"

type TransaksiCartPolisService interface {
	CreateTransaksiCartPolis(cartArr *[]models.TransaksiCart, carArr *[]models.TransaksiPolis) error
}
