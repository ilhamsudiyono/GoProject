package main

import (
	"ProjectFirst/config"
	"ProjectFirst/controller"
	"ProjectFirst/dao"
	"ProjectFirst/services"

	"gorm.io/gorm"
)

func main() {

	// init echo framework
	e := config.InitWeb()

	// init db and inject to dao and service
	g := initDb()
	dao.SetDao(g)
	services.SetService(g)

	// set jwt
	jwtGroup := config.SetJwt(e)

	// set controllers
	controller.SetInit(e)
	controller.SetUser(jwtGroup, e)
	controller.SetProduct(jwtGroup)
	controller.SetBp(jwtGroup)
	controller.SetTertanggung(jwtGroup)
	controller.SetPenerimaManfaat(jwtGroup)
	controller.SetTertanggungPm(jwtGroup)

	// start server
	e.Logger.Fatal(e.Start(":1233"))
}

func initDb() *gorm.DB {
	g, err := config.Conn()
	if err == nil {
		config.MigrateSchema(g)
		return g
	}
	panic(err)
}
