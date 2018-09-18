package model

import (
	"time"

	"github.com/lib/pq"
)

type Option struct {
	Name      string
	Position  int
	Product   Product
	Values    pq.StringArray `gorm:"type:varchar(100)[]"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
