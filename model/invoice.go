package model

type Invoice struct {
	ID             int     `json:"id"`
	Items           []int   `json:"items" db:"item_id"`
	Total          float32 `json:"total"`
	DateSubmmitted string  `json:"dateSubmitted"`
}
