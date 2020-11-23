package controller

import (
	"ProjectFirst/config"
	model "ProjectFirst/models"
	response "ProjectFirst/responses"
	service "ProjectFirst/services"
	"time"

	"github.com/labstack/echo"
)

var userService service.UserService = service.UserServiceImp{}

func SetUser(c *echo.Group, e *echo.Echo) {
	e.POST("/user", createUser)
	e.POST("/api/login", login)
	c.GET("/user/:id", getUserById)

}

func createUser(c echo.Context) (e error) {
	defer config.CatchError(&e)

	currentTime := time.Now()

	data := new(model.Users)

	if err := c.Bind(data); err != nil {
		return resErr(c, err)
	}

	result, err := userService.CreateUser(data)

	resPost := &response.UserResponsePostParent{response.BaseResponse{Status: "Created", StatusCode: 201, Message: "Data berhasil ditambah", Time: currentTime.String()},
		[]response.UserResponsePost{response.UserResponsePost{Email: result.Email, Gender: result.Gender}}}

	if err == nil {
		return respPost(c, resPost)
	}

	return resErr(c, err)
}

func getUserById(c echo.Context) (e error) {
	defer config.CatchError(&e)

	currentTime := time.Now()

	id := c.Param("id")

	var result, err = userService.GetUserById(id)

	resGet := &response.UserResponseGetIdParent{response.BaseResponse{Status: "OK", StatusCode: 200, Message: "Data berhasil ditampilkan", Time: currentTime.String()},
		[]response.UserResponseGetId{response.UserResponseGetId{Email: result.Email, Gender: result.Gender, RoleId: result.RoleId, BaseModel: result.BaseModel}}}

	if err == nil {
		return res(c, resGet)
	}
	return resErr(c, err)
}

func login(c echo.Context) (e error) {
	defer config.CatchError(&e)

	currentTime := time.Now()

	data := new(model.Users)

	if err := c.Bind(data); err != nil {
		return resErr(c, err)
	}

	var result, err = userService.Login(data.Email, data.Pwd)

	resPost := &response.UserResponseLoginParent{response.BaseResponse{Status: "OK", StatusCode: 200, Message: "Data berhasil login", Time: currentTime.String()},
		[]response.UserResponseLogin{response.UserResponseLogin{Email: result.Email, Gender: result.Gender, RoleId: result.RoleId, Token: *result.Token}}}

	if err == nil {
		return res(c, resPost)
	}

	return resErr(c, err)
}
