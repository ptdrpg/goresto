package repository

import (
	"errors"

	"github.com/ptdrpg/resto/entity"
)

func(r *Repository) VerifToken(username string) error {
	var staff entity.Staff
	if findStaffErr := r.DB.Where("username = ?", username).First(&staff).Error; findStaffErr != nil {
		return findStaffErr
	}

	if staff.Role != "admin" {
		return errors.New("not autorized to do this actions")
	}

	return nil
}