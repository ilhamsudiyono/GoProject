package config

import (
	model "ProjectFirst/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin2020"
	password = "admin123"
	dbname   = "db2020"
)

func Conn() (*gorm.DB, error) {
	config := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil

}

var tables = []interface{}{
	&model.Roles{},
	&model.PaymentType{},
	&model.Users{},
	&model.Product{},
	&model.ProductDtl{},
	&model.BusinessPartner{},
	&model.Tertanggung{},
	&model.PenerimaManfaat{},
	&model.TransaksiCart{},
	&model.TransaksiPolis{},
}

func MigrateSchema(db *gorm.DB) {
	db.AutoMigrate(tables...)
}
