package model

type Item struct {
	ID          int     `json:"id" db:"id"`
	Item        string  `json:"item"`
	Description string  `json:"description" db:"description"`
	Amount      float32 `json:"amount" db:"amount"`
}
