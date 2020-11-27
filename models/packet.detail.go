package models

// declaration model orders
type PacketDtl struct {
	PacketId    string `json:"packetId" gorm:"column:packet_id"`
	Packet      Packet `gorm:"foreignKey:PacketId"`
	NamaPaket   string `json:"namaPaket" gorm:"column:nama_paket"`
	Deskripsi1  string `json:"deskripsi1" gorm:"column:deskripsi_1"`
	Deskripsi2  string `json:"deskripsi2" gorm:"column:deskripsi_2"`
	JangkaWaktu string `json:"jangkaWaktu" gorm:"column:jangka_waktu"`
	Currency    string `json:"currency" gorm:"column:currency"`
	Tsi         int64  `json:"tsi" gorm:"column:total_sum_insured"`
	Premi       int64  `json:"premi" gorm:"column:premi"`
	BaseModel
}

// initiate table r_orders
func (PacketDtl) TableName() string {
	return "m_packet_detail"
}
