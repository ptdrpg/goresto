package repository

import "github.com/ptdrpg/resto/entity"

func (r *Repository) ListOrder() ([]entity.Order, error) {
	var orders []entity.Order
	orders = append(orders, entity.Order{ID: 5})
	return orders, nil
}
