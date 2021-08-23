package model

type Bill struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Owner        Owner  `json:"owner" db:"owner_id"`
	LastEdit     string `json:"last_edit" db:"lastEdit"`
	TemplateCode int    `json:"template_code" db:"template_code"`
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
	Number                int     `json:"number" db:"number_bid"`
	Owner                 Owner   `json:"owner" db:"owner"`
	SpecificationStimates string  `json:"specificationstimates" db:"specifications_stimates"`
	NotIncluded           string  `json:"not_included" db:"not_included"`
	TotalSum              float32 `json:"totalsum" db:"totalSum"`
	WithdrawnDays         int     `json:"withdrawndays" db:"withdrawn_days"`
	WithdrawnDate         string  `json:"withdrawndate" db:"withdrawn_date"`
	LastEdit              string  `db:"lastEdit"`
}

// BillJionBid La estructura que se espera
// sea devuelta para hacer el update
type BillJionBid struct {
	ID                    int     `json:"id" db:"id"`
	ID_BID                int     `json:"id_bid" db:"id_bid"`
	Template_code         int     `json:"template_code" db:"template_code"`
	Name                  string  `json:"name" db:"name"`
	Description           string  `json:"description" db:"description"`
	Number                int     `json:"number" db:"number_bid"`
	Owner                 Owner   `json:"owner" db:"owner"`
	SpecificationStimates string  `json:"specificationstimates" db:"specifications_stimates"`
	NotIncluded           string  `json:"not_included" db:"not_included"`
	TotalSum              float32 `json:"totalsum" db:"totalSum"`
	WithdrawnDays         int     `json:"withdrawndays" db:"withdrawn_days"`
	WithdrawnDate         string  `json:"withdrawndate" db:"withdrawn_date"`
	LastEdit              string  `json:"last_edit" db:"lastEdit"`
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

// BillJoinInvoice La estructura que se espera
// sea devuelta para hacer el update
type BillJoinInvoice struct {
	ID             int     `json:"id" db:"id_bill"`
	InvoiceID      int     `json:"invoice_id" db:"invoice_id"`
	Template_code  int     `json:"template_code" db:"template_code"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Number         int     `json:"number" db:"number_inv"`
	Owner          Owner   `json:"owner"`
	Item           []Item  `json:"item" db:"item"`
	Total          float32 `json:"total"`
	DateSubmmitted string  `json:"date_submmitted" db:"date_submmitted"`
	LastEdit       string  `json:"last_edit" db:"lastEdit"`
}
