package model

type Bill struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Owner       Owner    `json:"owner" db:"owner_id"`
	Template    Template `json:"template" db:"template_code"`
	LastEdit    string   `json:"last_edit" db:"lastEdit"`
}

type BillRequestGet struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Template_code int    `json:"template_code"`
	Description   string `json:"description"`
	LastEdit      string `json:"last_edit" db:"lastEdit"`
}

type BillBid struct {
	ID                    int     `json:"id"`
	Template_code         int     `json:"template_code"`
	Name                  string  `json:"name"`
	Description           string  `json:"description"`
	Number                int     `json:"number"`
	Owner                 Owner   `json:"owner"`
	SpecificationStimates string  `json:"specificationstimates"`
	NotIncluded           string  `json:"not_included"`
	TotalSum              float32 `json:"totalsum"`
	WithdrawnDays         int     `json:"withdrawndays"`
	WithdrawnDate         string  `json:"withdrawndate"`
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
