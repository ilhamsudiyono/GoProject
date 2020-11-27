package models

// declaration model orders
type Product struct {
	NamaProduk         string `json:"namaProduk" gorm:"column:nama_produk"`
	DeskripsiProduct   string `json:"deskripsiProduct" gorm:"column:deskripsi_produk"`
	DeskripsiGantiRugi string `json:"deskripsiGantiRugi" gorm:"column:deskripsi_ganti_rugi"`
	DeskripiTambahan   string `json:"deskripsiTambahan" gorm:"column:deskripsi_tambahan"`
	Pengguna           string `json:"pengguna" gorm:"column:pengguna"`
	SyaratKetentuan    string `json:"syarat&ketentuan" gorm:"column:syarat_dan_ketentuan"`
	CabangDefault      string `json:"cabangDefault" gorm:"column:cabang_default"`
	BaseModel
}

// initiate table r_orders
func (Product) TableName() string {
	return "m_product"
}
