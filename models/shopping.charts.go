package model

type ShoppingCharts struct {
	UserId   int32 `json:"userId" gorm:"column:user_id"`
	Users    Users `gorm:"foreignKey:UserId"`
	BookId   int32 `json:"bookId" gorm:"column:book_id"`
	Books    Books `gorm:"foreignKey:BookId"`
	Quantity int32 `json:"quantity" gorm:"column:quantity"`
	BaseModel
}

func (ShoppingCharts) TableName() string {
	return "r_shopping_charts"
}
