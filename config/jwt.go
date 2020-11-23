package config

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const secret = "secret"

func SetJwt(e *echo.Echo) *echo.Group {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	jwtGroup := e.Group("/api")
	jwtGroup.Use(middleware.JWT([]byte(secret)))
	return jwtGroup
}

func GenerateToken(username string) (string, error) {

	// create token
	token := jwt.New(jwt.SigningMethodHS256)

	//set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() //sejam

	t, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return t, nil
}
