package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Orders struct {
	OrderId         string         `json:"orderId" gorm:"primaryKey"`
	ShoppingChartId int32          `json:"shoppingChartId" gorm:"column:shopping_chart_id"`
	ShoppingCharts  ShoppingCharts `gorm:"foreignKey:ShoppingChartId"`
	Quantity        int32          `json:"quantity" gorm:"column:quantity"`
	PaymentTypeId   int32          `json:"paymentTypeId" gorm:"column:payment_type_id"`
	PaymentType     PaymentType    `gorm:"foreignKey:PaymentTypeId"`
	Currency        string         `json:"currency" gorm:"column:currency"`
	TotalPayment    uint32         `json:"totalPayment" gorm:"column:total_payment"`
	OrderDate       *Timestamp     `json:"createdDate" gorm:"type:timestamp without time zone;"`
	Information     string         `json:"information" gorm:"column:information"`
}

func (order *Orders) BeforeCreate(tx *gorm.DB) error {
	id := uuid.New()
	order.OrderId = id.String()
	return nil
}

func (Orders) TableName() string {
	return "r_orders"
}
