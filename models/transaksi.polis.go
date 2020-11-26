package models

// declaration model shoppingcharts
type TransaksiPolis struct {
	TransaksiCartId string        `json:"transaksiCartId" gorm:"column:transaksi_cart_id"`
	TransaksiCart   TransaksiCart `gorm:"foreignKey:TransaksiCartId"`
	ProductId       string        `json:"productId" gorm:"column:product_id"`
	Product         Product       `gorm:"foreignKey:ProductId"`
	ProductDtlId    string        `json:"productDtlId" gorm:"column:product_detail_id"`
	ProductDtl      ProductDtl    `gorm:"foreignKey:ProductDtlId"`

	JwpAwal     Timestamp `json:"jwpAwal" gorm:"column:jwp_awal"`
	JwpAkhir    Timestamp `json:"jwpAkhir" gorm:"column:jwp_akhir"`
	JsonData    string    `json:"jsonData" gorm:"column:json_data"`
	Currency    string    `json:"currency" gorm:"column:currency"`
	Tsi         float64   `json:"tsi" gorm:"column:tsi"`
	Premi       float64   `json:"premi" gorm:"column:premi"`
	NoPolis     string    `json:"noPolis" gorm:"column:no_polis"`
	IsClaimable bool      `json:"isClaimable" gorm:"column:is_claimable; default:false"`
	BaseModel
}

// initiate table r_shopping_charts
func (TransaksiPolis) TableName() string {
	return "r_transaksi_polis"
}
