package main

import (
	"coconut/db"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	defer db.GetDB().Close()

	router := drawRoutes()
	router.Run(":9876")
}
