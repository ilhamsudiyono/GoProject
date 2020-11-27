package models

import "ProjectFirst/helper"

// declaration model shoppingcharts
type TransaksiCart struct {
	BusinessPartnerId  string          `json:"businessPartnerId" gorm:"column:bp_id"`
	BusinessPartner    BusinessPartner `gorm:"foreignKey:BusinessPartnerId"`
	PaymentTypeId      string          `json:"paymentTypeId" gorm:"column:payment_type_id"`
	PaymentType        PaymentType     `gorm:"foreignKey:PaymentTypeId"`
	NoPengajuan        string          `json:"noPengajuanId" gorm:"column:no_pengajuan_id"`
	NoTransaksi        string          `json:"noTransaksiId" gorm:"column:no_transaksi_id"`
	TglTransaksi       helper.Date     `json:"tglTransaksi" gorm:"column:tgl_transaksi"`
	Currency           string          `json:"currency" gorm:"column:currency"`
	TotalTagihan       int64           `json:"totalTagihan" gorm:"column:total_tagihan"`
	NoReferensi        string          `json:"noReferensi" gorm:"column:no_referensi"`
	Keterangan         string          `json:"keterangan" gorm:"column:keterangan"`
	StatusPembayaran   int             `json:"statusPembayaran" gorm:"column:status_pembayaran"`
	TglPembayaran      helper.Date     `json:"tglPembayaran" gorm:"column:tgl_pembayaran"`
	PaymentExpiredDate helper.Date     `json:"paymentExpiredDate" gorm:"column:payment_expired_date"`
	StatusIssuedPolis  int             `json:"statusIssuedPolis" gorm:"column:status_issued_polis"`
	BaseModel
}

// initiate table r_shopping_charts
func (TransaksiCart) TableName() string {
	return "r_transaksi_cart"
}
