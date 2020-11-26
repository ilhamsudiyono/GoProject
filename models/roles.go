package models

// declaration model roles
type Roles struct {
	BaseModel
	RoleName string `json:"roleName" gorm:"column:role_name"`
	RoleCode string `json:"roleCode" gorm:"unique;column:role_code"`
}

// initiate table m_roles
func (Roles) TableName() string {
	return "m_roles"
}
