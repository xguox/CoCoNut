package main

import (
	. "github.com/xguox/coconut/config"
	"github.com/xguox/coconut/db"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// main =. =
func main() {
	defer db.GetDB().Close()
	router := drawRoutes()
	router.Run(Conf.Server.Port)
}
