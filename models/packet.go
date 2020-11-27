package models

// declaration model orders
type Packet struct {
	ProductDtlId string     `json:"productDtlId" gorm:"column:product_dtl_id"`
	ProductDtl   ProductDtl `gorm:"foreignKey:ProductDtlId"`
	NamaPaket    string     `json:"namaPaket" gorm:"column:nama_paket"`
	Deskripsi1   string     `json:"deskripsi1" gorm:"column:deskripsi_1"`
	Deskripsi2   string     `json:"deskripsi2" gorm:"column:deskripsi_2"`
	BaseModel
}

// initiate table r_orders
func (Packet) TableName() string {
	return "m_packet"
}
