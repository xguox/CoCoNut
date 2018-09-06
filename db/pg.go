package db

import (
	"fmt"
	"os/exec"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var PG *gorm.DB

func init() {
	var err error
	PG, err = gorm.Open("postgres", "user=postgres dbname=coconut_development sslmode=disable")
	if err != nil {
		panic(err)
	}
	PG.LogMode(true)
}

func TestDBInit() *gorm.DB {
	testDB, err := gorm.Open("postgres", "user=postgres dbname=coconut_test sslmode=disable")
	if err != nil {
		fmt.Println("db err: ", err)
	}
	testDB.LogMode(true)
	PG = testDB
	return PG
}

func ResetTestDB(testDB *gorm.DB) error {
	testDB.Close()

	cmd := exec.Command("RAILS_ENV=test rake db:drop && RAILS_ENV=test rake db:create && RAILS_ENV=test rake db:migrate")
	_, err := cmd.Output()

	return err
}

func GetDB() *gorm.DB {
	return PG
}
