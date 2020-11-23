package config

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func InitWeb() *echo.Echo {
	e := echo.New()
	e.Use(func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			log.Println("##########", req.Header.Get("Authorization"), "#########")
			if check := validateToken(req.Header.Get("Authorization")); !check {

				return c.JSON(http.StatusUnauthorized, "Invalid Token")

				log.Println("XXXXXXXXXXX")
			}

			return hf(c)
		}
	})
	return e
}

func validateToken(token string) bool {
	defer CatchErrorGeneral()
	log.Println("=====", token, "=====")
	return true

}
