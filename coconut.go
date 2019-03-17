package main

import (
	. "github.com/xguox/coconut/config"
	"github.com/xguox/coconut/db"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// @title Go Shopping Gin API
// @version 1.0
// @description e-commerce site written by go
// @termsOfService https://github.com/xguox/coconut
// @license.name MIT

func main() {
	defer db.GetDB().Close()
	router := drawRoutes()
	router.Run(Conf.Server.Port)
}
