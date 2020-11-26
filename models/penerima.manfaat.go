package models

// declaration model orders
type PenerimaManfaat struct {
	BusinessPartnerId string          `json:"businessPartnerId" gorm:"column:bp_id"`
	BusinessPartner   BusinessPartner `gorm:"foreignKey:BusinessPartnerId"`

	Nama        string `json:"nama" gorm:"column:nama"`
	Email       string `json:"email" gorm:"column:email"`
	NoTelp      string `json:"noTelp" gorm:"column:no_telp"`
	NoIdentitas int64  `json:"noIdentitas" gorm:"column:no_identitas"`
	IsVerified  bool   `json:"isVerfied" gorm:"default:false"`
	BaseModel
}

// initiate table r_orders
func (PenerimaManfaat) TableName() string {
	return "m_penerima_manfaat"
}
