package controller

import (
	"github.com/ptdrpg/resto/repository"
	"gorm.io/gorm"
)

type Controller struct {
	DB *gorm.DB
	R  *repository.Repository
}

func NewController(db *gorm.DB, r *repository.Repository) *Controller {
	return &Controller{
		DB: db,
		R:  r,
	}
}
