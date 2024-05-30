package entity

// type TicketItem struct {
// 	ID       uint `gorm:"primary_key"`
// 	TicketID uint
// 	Ticket   Ticket `gorm:"foreignKey:TicketID;references:ID"`
// 	ItemID   uint
// 	Item     Item `gorm:"foreignKey:ItemID;references:ID"`
// }

type ItemCount struct {
	ID       uint   `gorm:"primary_key"`
	ItemID   uint   `json:"-"`
	Item     Item   `json:"item" gorm:"-"`
	TicketID uint   `json:"-"`
	Ticket   Ticket `json:"ticket" gorm:"-"`
	Count    int    `json:"count"`
}
