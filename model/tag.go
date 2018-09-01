package model

import "time"

type Tag struct {
	ID            int
	Name          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	TaggingsCount uint
}

type Tagging struct {
	Tag     Tag
	Product Product
}
