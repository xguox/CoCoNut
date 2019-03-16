package db

import (
	. "github.com/xguox/coconut/config"
	"fmt"
	"os/exec"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	dbArgs := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=90s&parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci", Conf.Database.Username, Conf.Database.Pwd, Conf.Database.Host, Conf.Database.Port, Conf.Database.Basename)

	var err error
	DB, err = gorm.Open("mysql", dbArgs)

	if err != nil {
		panic(err)
	}
	DB.LogMode(true)
}

func TestDBInit() *gorm.DB {
	dbArgs := fmt.Sprintf("%s:%s@tcp(%s:%d)/coconut_test?timeout=90s&charset=utf8mb4&collation=utf8mb4_unicode_ci", Conf.Database.Username, Conf.Database.Pwd, Conf.Database.Host, Conf.Database.Port, Conf.Database.Basename)

	testDB, err := gorm.Open("mysql", dbArgs)
	if err != nil {
		fmt.Println("db err: ", err)
	}
	testDB.LogMode(true)
	DB = testDB
	return DB
}

func ResetTestDB(testDB *gorm.DB) error {
	testDB.Close()

	cmd := exec.Command("RAILS_ENV=test rake db:drop && RAILS_ENV=test rake db:create && RAILS_ENV=test rake db:migrate")
	_, err := cmd.Output()

	return err
}

func GetDB() *gorm.DB {
	return DB
}
