package models

// declaration model orders
type ProductDtl struct {
	ProductId  string  `json:"productId" gorm:"column:product_id"`
	Product    Product `gorm:"foreignKey:ProductId"`
	Nama       string  `json:"nama" gorm:"column:nama"`
	Deskripsi1 string  `json:"deskripsi1" gorm:"column:deskripsi_1"`
	Deskripsi2 string  `json:"deskripsi2" gorm:"column:deskripsi_2"`
	BaseModel
}

// initiate table r_orders
func (ProductDtl) TableName() string {
	return "m_product_detail"
}
