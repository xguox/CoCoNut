package model

import (
	"coconut/db"
)

func SaveData(data interface{}) error {
	err := db.PG.Save(data).Error
	return err
}
