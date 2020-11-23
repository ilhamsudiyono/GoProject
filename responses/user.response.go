package response

import (
	model "ProjectFirst/models"
)

// UserResponsePostParent ...
type UserResponsePostParent struct {
	BaseResponse
	UserResponsePost []UserResponsePost `json:"body"`
}

// UserResponsePost ...
type UserResponsePost struct {
	Email  string `json:"email"`
	Gender string `json:"gender"`
}

func NewUser(m *model.Users) *UserResponsePost {
	return &UserResponsePost{

		Email:  m.Email,
		Gender: m.Gender,
	}
}

// UserResponsePostParent ...
type UserResponseLoginParent struct {
	BaseResponse
	UserResponseLogin []UserResponseLogin `json:"payload"`
}

// UserResponsePost ...
type UserResponseLogin struct {
	Email  string `json:"email"`
	Gender string `json:"gender"`
	RoleId string `json:"roleId"`
	Token  string `json:"token"`
}

func PostUserLogin(m *model.Users) *UserResponseLogin {
	return &UserResponseLogin{
		Email:  m.Email,
		Gender: m.Gender,
		RoleId: m.RoleId,
		Token:  *m.Token,
	}
}
