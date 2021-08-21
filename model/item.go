package model

type Item struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Amount      float32 `json:"amount"`
}
