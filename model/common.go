package model

import (
	"github.com/xguox/coconut/db"
)

func SaveData(data interface{}) error {
	err := db.GetDB().Save(data).Error
	return err
}
