package model

type Roles struct {
	BaseModel
	RoleName string `json:"roleName" gorm:"column:role_name"`
	RoleCode string `json:"roleCode" gorm:"unique;column:role_code"`
}

func (Roles) TableName() string {
	return "m_roles"
}
