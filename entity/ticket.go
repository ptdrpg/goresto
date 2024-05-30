package entity

type Ticket struct {
	ID         uint        `gorm:"primaryKey" json:"id"`
	CustomerID int         `json:"customer_id"`
	Delivery   bool        `json:"delivery"`
	Total      int         `json:"total"`
	Items      []ItemCount `json:"-"`
	Date       string      `json:"date"`
	Updated_at string      `json:"updated_at"`
}
