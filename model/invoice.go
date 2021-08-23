package model

type Invoice struct {
	ID             int     `json:"id"`
	Number         int     `json:"number"`
	Item           []int   `json:"item" db:"item_id"`
	Total          float32 `json:"total"`
	DateSubmmitted string  `json:"date_submmitted"`
}
