package model

import "time"

type Tag struct {
	ID            uint
	Name          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	TaggingsCount uint
	Products      []Product `gorm:"many2many:taggings;"`
}

type Tagging struct {
	ID        uint
	TagID     uint
	ProductID uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
