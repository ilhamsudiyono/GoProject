package dao

import (
	"gorm.io/gorm"
)

// initiate g from datatype *gorm.DB
var g *gorm.DB

// create function for setting Dao
func SetDao(gDB *gorm.DB) {
	g = gDB
}
