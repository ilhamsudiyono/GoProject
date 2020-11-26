package responses

import (
	"ProjectFirst/helper"
	"ProjectFirst/models"
)

// UserResponsePostParent ...
type TertanggungResponsePostParent struct {
	BaseResponse
	TertangungResponsePost []TertangungResponsePost `json:"body"`
}

// UserResponsePost ...
type TertangungResponsePost struct {
	UserId            string      `json:"userId"`
	BusinessPartnerId string      `json:"businessPartnerId"`
	ProductId         string      `json:"productId"`
	Nama              string      `json:"nama"`
	NoIdentitas       int64       `json:"noIdentitas"`
	Email             string      `json:"email"`
	TglLahir          helper.Date `json:"tglLahir"`
	IsVerified        bool        `json:"IsVerified"`
}

func NewTertanggung(m *models.Tertanggung) *TertangungResponsePost {
	return &TertangungResponsePost{
		UserId:            m.UserId,
		BusinessPartnerId: m.BusinessPartnerId,
		ProductId:         m.ProductId,
		Nama:              m.Nama,
		NoIdentitas:       m.NoIdentitas,
		Email:             m.Email,
		TglLahir:          m.TglLahir,
		IsVerified:        m.IsVerified,
	}
}
