package models

type Product struct {
	ID int `json:"id"`

	Name string `json:"name"`

	Qty int `json:"qty"`

	LastUpdated string `json:"last_updated"`
}
