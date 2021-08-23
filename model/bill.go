package model

type Bill struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Owner        Owner  `json:"owner" db:"owner_id"`
	LastEdit     string `json:"last_edit" db:"lastEdit"`
	TemplateCode int    `json:"tempalte_code"`
}

type BillRequestGet struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Template_code int    `json:"template_code"`
	Description   string `json:"description"`
	LastEdit      string `json:"last_edit" db:"lastEdit"`
}

type BillBid struct {
	ID                    int     `json:"id" db:"id"`
	Template_code         int     `json:"template_code" db:"template_code"`
	Name                  string  `json:"name" db:"name"`
	Description           string  `json:"description" db:"description"`
	Number                int     `json:"number" db:"number_id"`
	Owner                 Owner   `json:"owner" db:"owner"`
	SpecificationStimates string  `json:"specificationstimates" db:"specifications_stimates"`
	NotIncluded           string  `json:"not_included" db:"not_included"`
	TotalSum              float32 `json:"totalsum" db:"totalSum"`
	WithdrawnDays         int     `json:"withdrawndays" db:"withdrawn_days"`
	WithdrawnDate         string  `json:"withdrawndate" db:"withdrawn_date"`
}

type BillInvoice struct {
	ID             int     `json:"id"`
	Template_code  int     `json:"template_code"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Number         int     `json:"number"`
	Owner          Owner   `json:"owner"`
	Item           []Item  `json:"item" db:"item_id"`
	Total          float32 `json:"total"`
	DateSubmmitted string  `json:"date_submmitted"`
}
