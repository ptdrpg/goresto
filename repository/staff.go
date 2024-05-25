package repository

import (
	"errors"

	"github.com/ptdrpg/resto/entity"
)

func (r *Repository) FindAllStaff() ([]entity.Staff, error) {
	var staffs []entity.Staff
	if err := r.DB.Model(&entity.Staff{}).Find(&staffs).Error; err != nil {
		return []entity.Staff{}, err
	}

	return staffs, nil
}

func (r *Repository) FindStaffById(id int) (staff entity.Staff, err error) {
	var findStaff entity.Staff
	result := r.DB.Find(&findStaff, id)
	if result != nil {
		return findStaff, nil
	} else {
		return findStaff, errors.New("staff not found")
	}
}

func (r *Repository) CreateStaff(staff *entity.Staff) error {
	if err := r.DB.Create(staff).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateStaff(staff *entity.Staff) error {
	if err := r.DB.Model(staff).Updates(staff).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteStaff(id int) error {
	var staff entity.Staff
	if err := r.DB.Where("id = ?", id).Delete(&staff).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdatePassword(pass *entity.Staff) error {
	if err := r.DB.Model(pass).Updates(pass).Error; err != nil {
		return err
	}

	return nil
}
