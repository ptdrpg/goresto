package repository

import (
	"errors"

	"github.com/ptdrpg/resto/entity"
)

func (r *Repository) FindAllTicket() ([]entity.Ticket, error) {
	var tickets []entity.Ticket
	if err := r.DB.Model(&entity.Ticket{}).Find(&tickets).Error; err != nil {
		return []entity.Ticket{}, nil
	}

	return tickets, nil
}

func (r *Repository) GetAllItemCount() ([]entity.ItemCount, error) {
	var count []entity.ItemCount
	if err := r.DB.Model(&entity.ItemCount{}).Find(&count).Error; err != nil {
		return []entity.ItemCount{}, nil
	}

	return count, nil
}

func (r *Repository) FindTicketById(id int) (entity.Ticket, error) {
	var ticket entity.Ticket
	result := r.DB.Where("id = ?", id).Find(&ticket)
	if result != nil {
		return ticket, nil
	} else {
		return ticket, errors.New("ticket not found")
	}
}

func (r *Repository) CreateTicket(ticket *entity.Ticket) error {
	if err := r.DB.Create(ticket).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateTicket(ticket *entity.Ticket) error {
	if err := r.DB.Model(ticket).Updates(ticket).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteTicket(ticket *entity.Ticket) error {
	// var ticket entity.Ticket
	// r.DB.Where("id = ?", id).First(&ticket)

	var itemCount entity.ItemCount
	if err := r.DB.Where("ticket_id = ?", ticket.ID).Delete(&itemCount).Error; err != nil {
		return err
	}

	// err := r.DB.Raw("DELETE FROM tickets WHERE id = ?", id)
	if err := r.DB.Delete(ticket).Error; err != nil {
		return err
	}

	return nil
}
