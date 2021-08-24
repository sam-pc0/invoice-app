package model

type Invoice struct {
	ID             int     `json:"id"`
	Item           []int   `json:"item" db:"item_id"`
	Total          float32 `json:"total"`
	DateSubmmitted string  `json:"dateSubmitted"`
}
