package entity

import "gorm.io/gorm"

type Entity struct {
	DB *gorm.DB
}

func NewEntity(db *gorm.DB) *Entity {
	return &Entity{DB: db}
}
