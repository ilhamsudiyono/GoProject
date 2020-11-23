package model

type PaymentType struct {
	BaseModel
	Name        string `json:"name" gorm:"column:payment_name"`
	Description string `json:"description" gorm:"column:description"`
}

func (PaymentType) TableName() string {
	return "m_payment_type"
}
