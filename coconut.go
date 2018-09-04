package main

import (
	"coconut/db"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	defer db.PG.Close()

	router := drawRoutes()
	router.Run(":9876")
}
