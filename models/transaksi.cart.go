package models

// declaration model shoppingcharts
type TransaksiCart struct {
	PaymentTypeId      string      `json:"paymentTypeId" gorm:"column:payment_type_id"`
	PaymentType        PaymentType `gorm:"foreignKey:PaymentTypeId"`
	NoPengajuan        int64       `json:"noPengajuanId" gorm:"column:no_pengajuan_id"`
	NoTransaksi        int64       `json:"noTransaksiId" gorm:"column:no_transaksi_id"`
	TglTransaksi       Timestamp   `json:"tglTransaksi" gorm:"column:tgl_transaksi"`
	Currency           string      `json:"currency" gorm:"column:currency"`
	TotalTagihan       float64     `json:"totalTagihan" gorm:"column:total_tagihan"`
	NoReferensi        int64       `json:"noReferensi" gorm:"column:no_referensi"`
	Keterangan         int64       `json:"keterangan" gorm:"column:keterangan"`
	StatusPembayaran   int         `json:"statusPembayaran" gorm:"column:status_pembayaran"`
	TglPembayaran      Timestamp   `json:"tglPembayaran" gorm:"column:tgl_pembayaran"`
	PaymentExpiredDate Timestamp   `json:"paymentExpiredDate" gorm:"column:payment_expired_date"`
	StatusIssuedPolis  Timestamp   `json:"statusIssuedPolis" gorm:"column:status_issued_polis"`
	BaseModel
}

// initiate table r_shopping_charts
func (TransaksiCart) TableName() string {
	return "r_transaksi_cart"
}
