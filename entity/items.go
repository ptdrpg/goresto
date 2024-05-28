package entity

type Items struct {
	ID         uint   `json:"id"`
	Label      string `json:"label"`
	Short_desc string `json:"short_desc"`
	Price      int    `json:"price"`
	Category   string `json:"category"`
}
