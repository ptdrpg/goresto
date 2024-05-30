package repository

import (
	"errors"

	"github.com/ptdrpg/resto/entity"
)

func (r *Repository) FindAllItems() ([]entity.Item, error) {
	var items []entity.Item
	if err := r.DB.Model(&entity.Item{}).Find(&items).Error; err != nil {
		return []entity.Item{}, nil
	}

	return items, nil
}

func (r *Repository) FindItemById(id int) (entity.Item, error) {
	var item entity.Item
	result := r.DB.Find(&item, id)
	if result != nil {
		return item, nil
	} else {
		return item, errors.New("item not found")
	}
}

func (r *Repository) CreateItems(item *entity.Item) error {
	err := r.DB.Create(item)
	if err != nil {
		return err.Error
	}

	return nil
}

func (r *Repository) UpdateItems(item *entity.Item) error {
	err := r.DB.Model(item).Updates(item)
	if err != nil {
		return err.Error
	}

	return nil
}

func (r *Repository) DeleteItems(id int) error {
	var item entity.Item
	if err := r.DB.Where("id = ?", id).Delete(&item).Error; err != nil {
		return err
	}

	return nil
}
