package responses

import (
	"ProjectFirst/helper"
	"ProjectFirst/models"
)

// UserResponsePostParent ...
type BpResponsePostParent struct {
	BaseResponse
	BpResponsePost []BpResponsePost `json:"body"`
}

// UserResponsePost ...
type BpResponsePost struct {
	UserId      string      `json:"userId"`
	Nama        string      `json:"nama"`
	NoIdentitas string      `json:"noIdentitas"`
	Email       string      `json:"email"`
	Alamat      string      `json:"alamat"`
	TglLahir    helper.Date `json:"tglLahir"`
	IsVerified  bool        `json:"IsVerified"`
}

func NewBp(m *models.BusinessPartner) *BpResponsePost {
	return &BpResponsePost{
		UserId:      m.UserId,
		Nama:        m.Nama,
		NoIdentitas: m.NoIdentitas,
		Email:       m.Email,
		Alamat:      m.Alamat,
		TglLahir:    m.TglLahir,
		IsVerified:  m.IsVerified,
	}
}
