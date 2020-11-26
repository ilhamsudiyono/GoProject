package models

import "ProjectFirst/helper"

// declaration model orders
type BusinessPartner struct {
	UserId      string      `gorm:"column:user_id"`
	Users       Users       `gorm:"foreignKey:UserId"`
	Nama        string      `gorm:"column:nama"`
	NoIdentitas string      `gorm:"column:no_identitas"`
	Email       string      `gorm:"column:email"`
	Alamat      string      `gorm:"column:alamat"`
	TglLahir    helper.Date `gorm:"column:tgl_lahir; type:date"`
	IsVerified  bool        `gorm:"default:true"`
	BaseModel
}

// initiate table r_orders
func (BusinessPartner) TableName() string {
	return "m_business_partner"
}
