package model

type Bill struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Owner       Owner    `json:"owner" db:"owner_id"`
	Template    Template `json:"template" db:"template_code"`
	LastEdit    string   `json:"last_edit" db:"lastEdit"`
}
