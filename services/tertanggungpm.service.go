package services

import "ProjectFirst/models"

type TertanggungPmService interface {
	CreateTertanggungPm(trArr *[]models.Tertanggung, pmArr *[]models.PenerimaManfaat) error
}
