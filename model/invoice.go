package model

type Invoice struct {
	ID             int     `json:"id"`
	Items          []int   `json:"items" db:"item_id"`
	DateSubmmitted string  `json:"dateSubmitted"`
}

type BillJoinInvoice struct {
	ID             int     `json:"id" db:"id"`
	Template_code  int     `json:"template_code" db:"template_code"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Owner          Owner   `json:"owner"`
	Items          []Item  `json:"items" db:"item"`
	DateSubmmitted string  `json:"dateSubmitted" db:"dateSubmmitted"`
	LastEdit       string  `json:"last_edit" db:"lastEdit"`
	Total          float32 `json:"total" db:"total"`
	SubTotal       float32 `json:"subTotal" db:"sub_total"`
	TaxRate        float32 `json:"taxRate" db:"tax_rate"`
	Tax            float32 `json:"tax" db:"tax"`
}
