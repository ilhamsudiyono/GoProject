package requests

import (
	"ProjectFirst/helper"
	"ProjectFirst/models"

	"github.com/labstack/echo"
)

// UserRequest ...
type BusinessPartnerRequest struct {
	UserId      string      `form:"userId" json:"userId" validate:"required"`
	Nama        string      `form:"nama" json:"nama" validate:"required"`
	NoIdentitas string      `form:"noIdentitas" json:"noIdentitas" validate:"required"`
	Email       string      `form:"email" json:"email" validate:"required"`
	Alamat      string      `form:"alamat" json:"alamat" validate:"required"`
	TglLahir    helper.Date `form:"tglLahir" json:"tglLahir" validate:"required"`
	IsVerified  bool        `form:"isVerified" json:"isVerified" validate:"required"`
}

// Convert from echo FormValue
func (b BusinessPartnerRequest) Convert(c echo.Context) *BusinessPartnerRequest {
	c.Bind(&b)
	return &b
}

// Model from request
func (b BusinessPartnerRequest) Model() *models.BusinessPartner {
	return &models.BusinessPartner{
		UserId:      b.UserId,
		Nama:        b.Nama,
		NoIdentitas: b.NoIdentitas,
		Email:       b.Email,
		Alamat:      b.Alamat,
		TglLahir:    b.TglLahir,
		IsVerified:  b.IsVerified,
	}
}
