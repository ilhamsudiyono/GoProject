package models

// declaration model PaymentType
type PaymentType struct {
	BaseModel
	Name        string `json:"name" gorm:"column:payment_name"`
	Description string `json:"description" gorm:"column:description"`
}

// initiate table m_payment_type
func (PaymentType) TableName() string {
	return "m_payment_type"
}
