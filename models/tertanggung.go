package models

import "ProjectFirst/helper"

// declaration model orders
type Tertanggung struct {
	UserId            string          `json:"userId" gorm:"column:user_id"`
	Users             Users           `gorm:"foreignKey:UserId"`
	BusinessPartnerId string          `json:"businessPartnerId" gorm:"column:bp_id"`
	BusinessPartner   BusinessPartner `gorm:"foreignKey:BusinessPartnerId"`
	ProductId         string          `json:"productId" gorm:"column:product_id"`
	Product           Product         `gorm:"foreignKey:ProductId"`

	Nama        string      `json:"nama" gorm:"column:nama"`
	Email       string      `json:"email" gorm:"column:email"`
	NoTelp      string      `json:"noTelp" gorm:"column:no_telp"`
	NoIdentitas int64       `json:"noIdentitas" gorm:"column:no_identitas"`
	TglLahir    helper.Date `json:"tglLahir" gorm:"column:tgl_lahir"`
	IsVerified  bool        `json:"isVerfied" gorm:"default:true"`
	BaseModel
}

// initiate table r_orders
func (Tertanggung) TableName() string {
	return "m_tertanggung"
}
