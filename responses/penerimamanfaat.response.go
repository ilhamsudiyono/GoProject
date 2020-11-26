package responses

import (
	"ProjectFirst/models"
)

// UserResponsePostParent ...
type PenerimaManfaatResponsePostParent struct {
	BaseResponse
	PenerimaManfaatResponsePost []PenerimaManfaatResponsePost `json:"body"`
}

// UserResponsePost ...
type PenerimaManfaatResponsePost struct {
	BusinessPartnerId string `json:"businessPartnerId"`
	Nama              string `json:"nama"`
	Email             string `json:"email"`
	NoTelp            string `json:"noTelp"`
	NoIdentitas       int64  `json:"noIdentitas"`
	IsVerified        bool   `json:"IsVerified"`
}

func NewPenerimaManfaat(m *models.PenerimaManfaat) *PenerimaManfaatResponsePost {
	return &PenerimaManfaatResponsePost{
		BusinessPartnerId: m.BusinessPartnerId,
		Nama:              m.Nama,
		Email:             m.Email,
		NoTelp:            m.NoTelp,
		NoIdentitas:       m.NoIdentitas,
		IsVerified:        m.IsVerified,
	}
}
