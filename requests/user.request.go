package requests

import (
	"ProjectFirst/models"

	"github.com/labstack/echo"
)

// UserRequest ...
type UserRequest struct {
	Email     string `form:"email" json:"email" validate:"required"`
	Password  string `form:"password" json:"password" validate:"required"`
	Gender    string `form:"gender" json:"gender" validate:"required"`
	RoleId    string `form:"roleId" json:"roleId" validate:"required"`
	IsActive  bool   `form:"isActive" json:"isActive"`
	CreatedBy string `form:"createdBy" json:"createdBy"`
}

// Convert from echo FormValue
func (u UserRequest) Convert(c echo.Context) *UserRequest {
	c.Bind(&u)
	return &u
}

// Model from request
func (u UserRequest) Model() *models.Users {
	return &models.Users{
		Email:  u.Email,
		Pwd:    u.Password,
		Gender: u.Gender,
		RoleId: u.RoleId,
	}
}
