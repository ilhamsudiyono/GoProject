package model

type Books struct {
	Title        string  `json:"title_book" gorm:"column:title_buku"`
	Noisbn       int64   `json:"noisbn" gorm:"column:noisbn"`
	Author       string  `json:"author" gorm:"column:author"`
	Publisher    string  `json:"publisher" gorm:"column:publisher"`
	Year         uint8   `json:"year" gorm:"column:year"`
	Stock        uint32  `json:"stock" gorm:"column:stock"`
	Cost         float64 `json:"cost" gorm:"column:cost"`
	SellingPrice float64 `json:"selling_price" gorm:"column:selling_price"`
	Discount     uint32  `json:"discount" gorm:"column:discount"`
	PpnPajak     uint32  `json:"ppn" gorm:"column:ppn"`
	BaseModel
}

func (Books) TableName() string {
	return "m_books"
}
