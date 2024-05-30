package repository

import "github.com/ptdrpg/resto/entity"

func (r *Repository) FindTicketItems(ticket_id uint) ([]entity.Item, error) {
	var items []entity.Item
	r.DB.InnerJoins("JOIN \"ticket_item\" ON \"ticket_item\".\"item_id\" = \"item\".\"id\"").
		InnerJoins("JOIN \"ticket\" ON \"ticket_item\".\"ticket_id\" = \"ticket\".\"id\"").
		Where("\"ticket\".\"id\" = ?", ticket_id).Find(&items)
	return items, nil
}

func (r *Repository) AppendItem2Ticket(ticket entity.Ticket, item entity.Item) error {
	if err := r.DB.Model(&ticket).Association("Items").Append(&item); err != nil {
		return err
	}
	return nil
}

func (r *Repository) CreateItemCount(itemC *entity.ItemCount) error {
	r.DB.Create(itemC)
	return nil
}
