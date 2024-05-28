package entity

type Customer struct {
	ID           uint   `gorm:"primary_key" json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone_number int    `json:"phone_number"`
	Address      string `json:"address"`
	Age          string `json:"age"`
	Gender       string `json:"gender"`
	Point        int    `json:"point"`
}
