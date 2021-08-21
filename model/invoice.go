package model

import "time"

type Invoice struct {
	ID             int       `json:"id"`
	TemplateCode   Template  `json:"template" db:"template_code"`
	Number         int       `json:"number"`
	Item           Item      `json:"item" db:"item_id"`
	Total          float32   `json:"total"`
	DateSubmmitted time.Time `json:"date_submmitted"`
}
