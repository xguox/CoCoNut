package main

import (
	. "coconut/config"
	"coconut/db"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// main =. =
func main() {
	defer db.GetDB().Close()
	router := drawRoutes()
	router.Run(Conf.Server.Port)
}
