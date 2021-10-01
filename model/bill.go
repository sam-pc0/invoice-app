package model

type Bill struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Owner        Owner  `json:"owner" db:"owner_id"`
	LastEdit     string `json:"last_edit" db:"lastEdit"`
	TemplateCode int    `json:"template_code" db:"template_code"`
	Total                 float32 `json:"total" db:"total"`
	SubTotal              float32 `json:"subTotal" db:"sub_total"`
	TaxRate               float32 `json:"taxRate" db:"tax_rate"`
	Tax                   float32 `json:"tax" db:"tax"`
}

type BillRequestGet struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Template_code int    `json:"template_code"`
	Description   string `json:"description"`
	LastEdit      string `json:"last_edit" db:"lastEdit"`
}
