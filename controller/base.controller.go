package controller

import (
	"ProjectFirst/helper"
	"net/http"

	"github.com/go-playground/validator"

	"github.com/labstack/echo"
)

func SetInit(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Rest API Started")
	})
}

func res(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}

func respPost(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusCreated, data)
}

func resErr(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, err.Error())
}

func resValErr(c echo.Context, err error) error {
	var errorMessages []string
	validationErrors := err.(validator.ValidationErrors)
	for _, v := range validationErrors {
		var message string
		switch v.Tag() {
		case "required":
			message = helper.ToSnakeCase(v.Field()) + " is required"
		case "email":
			message = helper.ToSnakeCase(v.Field()) + " must be an email"
		}
		errorMessages = append(errorMessages, message)
	}
	return c.JSON(http.StatusBadRequest, errorMessages)
}
