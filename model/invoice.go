package model

type Invoice struct {
	ID             int      `json:"id"`
	TemplateCode   Template `json:"template" db:"template_code"`
	Number         int      `json:"number"`
	Item           []int    `json:"item" db:"item_id"`
	Total          float32  `json:"total"`
	DateSubmmitted string   `json:"date_submmitted"`
}
