package models

type TransaksiCartPolisPojo struct {
	TransaksiCart  []TransaksiCart  `json:"transaksiCart"`
	TransaksiPolis []TransaksiPolis `json:"transaksiPolis"`
}
