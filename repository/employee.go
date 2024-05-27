package repository

import (
	"errors"

	"github.com/ptdrpg/resto/entity"
)

func (r *Repository) FindAllEmployee() ([]entity.Employee, error) {
	var employes []entity.Employee
	if err := r.DB.Model(&entity.Employee{}).Find(&employes).Error; err != nil {
		return []entity.Employee{}, nil
	}

	return employes, nil
}

func (r *Repository) FindEmployeeById(id int) (entity.Employee, error) {
	var findEmployee entity.Employee
	result := r.DB.Find(&findEmployee, id)
	if result != nil {
		return findEmployee, nil
	} else {
		return findEmployee, errors.New("staff not found")
	}
}

func (r *Repository) CreateEmployee(employee *entity.Employee) error {
	if err := r.DB.Create(employee).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateEmployee(employee *entity.Employee) error {
	if err := r.DB.Model(employee).Updates(employee).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteEmployee(id int) error {
	var employee entity.Employee
	if err := r.DB.Where("id = ?", id).Delete(&employee).Error; err != nil {
		return err
	}

	return nil
}
