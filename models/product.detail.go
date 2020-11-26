package models

// declaration model orders
type ProductDtl struct {
	Nama        string  `json:"nama" gorm:"column:nama"`
	JangkaWaktu int     `json:"jangkaWaktu" gorm:"column:jangka_waktu"`
	Tsi         float64 `json:"tsi" gorm:"column:tsi"`
	Premi       float64 `json:"premi" gorm:"column:premi"`
	BaseModel
}

// initiate table r_orders
func (ProductDtl) TableName() string {
	return "m_product_detail"
}
