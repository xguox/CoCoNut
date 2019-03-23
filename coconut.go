package main

import (
	_ "net/http/pprof"

	"github.com/DeanThompson/ginpprof"
	. "github.com/xguox/coconut/config"
	"github.com/xguox/coconut/db"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// @title Go Shopping Gin API
// @version 1.0
// @description e-commerce site written by go
// @termsOfService https://github.com/xguox/coconut
// @license.name MIT

// @contact.name XguoX
// @contact.url https://xguox.me

// @license.name MIT
// @host localhost:9876
// @BasePath /api/v1
func main() {
	defer db.GetDB().Close()
	router := drawRoutes()
	ginpprof.Wrap(router)

	router.Run(Conf.Server.Port)
}
