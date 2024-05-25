package repository

import (
	"errors"

	"github.com/ptdrpg/resto/entity"
)

func (u *Repository) FindAllUsers() ([]entity.Customer, error) {
	var users []entity.Customer
	if err := u.DB.Model(&entity.Customer{}).Find(&users).Error; err != nil {
		return []entity.Customer{}, err
	}

	return users, nil
}

func (u *Repository) FindUserById(id int) (user entity.Customer, err error) {
	var findUser entity.Customer
	result := u.DB.Find(&findUser, id)
	if result != nil {
		return findUser, nil
	} else {
		return findUser, errors.New("user not found")
	}
}

func (u *Repository) CreateUser(creatUserDto *entity.Customer) error {
	if err := u.DB.Create(creatUserDto).Error; err != nil {
		return err
	}
	return nil
}

func (u *Repository) UpdateUser(updateUserDto *entity.Customer) {
	var updateUser = entity.Customer{
		Name:         updateUserDto.Name,
		Email:        updateUserDto.Email,
		Phone_number: updateUserDto.Phone_number,
		Age:          updateUserDto.Age,
		Gender:       updateUserDto.Gender,
		ID:           updateUserDto.ID,
		Address:      updateUserDto.Address,
	}
	u.DB.Model(updateUserDto).Updates(updateUser)
}

func (u *Repository) DeleteUser(id int) error {
	var user entity.Customer
	if err := u.DB.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
