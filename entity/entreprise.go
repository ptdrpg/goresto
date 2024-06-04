package entity

type Entreprise struct {
	ID                 uint   `json:"id"`
	Entreprise_name    string `json:"entreprise_name"`
	Entreprise_address string `json:"entreprise_address"`
	Avatar             string `json:"avatar"`
	AdminID            uint   `json:"-"`
}
