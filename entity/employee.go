package entity

type Employee struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	CustomerID int    `json:"customer_id"`
	Job        string `json:"job"`
	Hire_date  string `json:"hire_date"`
	Avatar     string `json:"avatar"`
}
