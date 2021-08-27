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
	Owner                 Owner   `json:"owner" db:"owner"`
	SpecificationStimates string  `json:"specificationstimates" db:"specifications_stimates"`
	NotIncluded           string  `json:"notIncluded" db:"not_included"`
	TotalSum              float32 `json:"totalsum" db:"totalSum"`
	WithdrawnDays         int     `json:"withdrawnDays" db:"withdrawn_days"`
	WithdrawnDate         string  `json:"withdrawnDate" db:"withdrawn_date"`
	LastEdit              string  `db:"lastEdit"`
}

// BillJionBid La estructura que se espera
// sea devuelta para hacer el update
type BillJoinBid struct {
	ID                    int     `json:"id" db:"id"`
	ID_BID                int     `json:"id_bid" db:"id_bid"`
	Template_code         int     `json:"template_code" db:"template_code"`
	Name                  string  `json:"name" db:"name"`
	Description           string  `json:"description" db:"description"`
	Owner                 Owner   `json:"owner" db:"owner"`
	SpecificationStimates string  `json:"specificationstimates" db:"specifications_stimates"`
	NotIncluded           string  `json:"notIncluded" db:"not_included"`
	TotalSum              float32 `json:"totalsum" db:"totalSum"`
	WithdrawnDays         int     `json:"withdrawnDays" db:"withdrawn_days"`
	WithdrawnDate         string  `json:"withdrawnDate" db:"withdrawn_date"`
	LastEdit              string  `json:"last_edit" db:"lastEdit"`
}

type BillInvoice struct {
	ID             int     `json:"id"`
	Template_code  int     `json:"template_code"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Owner          Owner   `json:"owner"`
	Item           []Item  `json:"item"`
	Total          float32 `json:"total"`
	DateSubmmitted string  `json:"dateSubmitted"`
}

// BillJoinInvoice La estructura que se espera
// sea devuelta para hacer el update
type BillJoinInvoice struct {
	ID             int     `json:"id" db:"id"`
	Template_code  int     `json:"template_code" db:"template_code"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Owner          Owner   `json:"owner"`
	Items          []Item  `json:"items" db:"item"`
	Total          float32 `json:"total"`
	DateSubmmitted string  `json:"dateSubmitted" db:"date_submmitted"`
	LastEdit       string  `json:"last_edit" db:"lastEdit"`
}
