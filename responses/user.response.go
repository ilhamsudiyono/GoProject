package responses

import (
	"ProjectFirst/models"
	model "ProjectFirst/models"
	"time"
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
	models.BaseModel
	Email  string `json:"email"`
	Gender string `json:"gender"`
	RoleId string `json:"roleId"`
	Token  string `json:"token"`
}

func PostUserLogin(m *model.Users) *UserResponseLogin {
	return &UserResponseLogin{
		BaseModel: m.BaseModel,
		Email:     m.Email,
		Gender:    m.Gender,
		RoleId:    m.RoleId,
		Token:     *m.Token,
	}
}

// UserResponsePostParent ...
type UserResponseGetIdParent struct {
	BaseResponse
	UserResponseGetId []UserResponseGetId `json:"payload"`
}

// UserResponsePost ...
type UserResponseGetId struct {
	model.BaseModel
	Email  string `json:"email"`
	Gender string `json:"gender"`
	RoleId string `json:"roleId"`
}

func GetUserById(m *model.Users) *UserResponseGetIdParent {
	currentTime := time.Now()

	return &UserResponseGetIdParent{

		UserResponseGetId: []UserResponseGetId{UserResponseGetId{Email: m.Email, Gender: m.Gender, RoleId: m.RoleId, BaseModel: m.BaseModel}},
		BaseResponse:      BaseResponse{Status: "", StatusCode: 200, Message: "Data berhasil ditampilkan", Time: currentTime.String()},
	}
}
