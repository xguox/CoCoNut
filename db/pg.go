package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var PG *gorm.DB

func init() {
	var err error
	PG, err = gorm.Open("postgres", "user=postgres dbname=coconut_development sslmode=disable")
	PG.LogMode(true)
	if err != nil {
		panic(err)
	}
	// db.AutoMigrate(&topic{})
}
