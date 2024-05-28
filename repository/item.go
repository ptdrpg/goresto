package repository

import (
	"errors"

	"github.com/ptdrpg/resto/entity"
)

func (r *Repository) FindAllItems() ([]entity.Items, error) {
	var items []entity.Items
	if err := r.DB.Model(&entity.Items{}).Find(&items).Error;err != nil {
		return []entity.Items{}, nil
	}

	return items, nil
}

func (r *Repository) FindItemById(id int) (entity.Items, error) {
	var item entity.Items
	result := r.DB.Find(&item, id)
	if result != nil {
		return item, nil
	} else {
		return item, errors.New("item not found")
	}
}

func (r *Repository) CreateItems(item *entity.Items) (error) {
	err := r.DB.Create(item)
	if err != nil {
		return err.Error
	}

	return nil
}

func (r *Repository) UpdateItems(item *entity.Items) (error) {
	err := r.DB.Model(item).Updates(item)
	if err != nil {
		return err.Error
	}

	return nil
}

func (r *Repository) DeleteItems(id int) error {
	var item entity.Items
	if err := r.DB.Where("id = ?", id).Delete(&item).Error; err != nil {
		return err
	}

	return nil
}
