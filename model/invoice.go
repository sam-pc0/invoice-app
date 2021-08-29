package model

type Invoice struct {
	ID             int     `json:"id"`
	Items           []int   `json:"items" db:"item_id"`
	Total          float32 `json:"total"`
	DateSubmmitted string  `json:"dateSubmitted"`
}

type BillJoinInvoice struct {
	ID             int     `json:"id" db:"id"`
	Template_code  int     `json:"template_code" db:"template_code"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Owner          Owner   `json:"owner"`
	Items          []Item  `json:"items" db:"item"`
	Total          float32 `json:"total"`
	DateSubmmitted string  `json:"dateSubmitted" db:"dateSubmmitted"`
	LastEdit       string  `json:"last_edit" db:"lastEdit"`
}