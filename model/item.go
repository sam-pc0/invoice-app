package model

type Item struct {
	ID          int     `json:"id" db:"id"`
	Description string  `json:"description" db:"description"`
	Amount      float32 `json:"amount" db:"amount"`
}
