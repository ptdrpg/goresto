package entity

type Order struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}
