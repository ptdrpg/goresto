package entity

type Staff struct {
	ID           uint   `gorm:"primary_key" json:"id"`
	CustomerID   int    `json:"customer_id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Avatar       string `json:"avatar"`
	EntrepriseID int    `json:"-"`
	Role         string `json:"role"`
}
