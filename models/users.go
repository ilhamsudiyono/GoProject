package model

type Users struct {
	Email  string  `json:"email" gorm:"unique; column:email"`
	Pwd    string  `json:"pwd" gorm:"column:password"`
	Gender string  `json:"gender" gorm:"column:gender"`
	RoleId string  `json:"roleId" gorm:"column:role_id"`
	Roles  Roles   `gorm:"foreignKey:RoleId"`
	Token  *string `json:"token" gorm:"-"`
	BaseModel
}

func (Users) TableName() string {
	return "m_users"
}
