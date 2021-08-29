package model

type BidProposal struct {
	ID                    int     `json:"id"`
	SpecificationStimates string  `json:"specificationstimates"`
	NotIncluded           string  `json:"notIncluded"`
	TotalSum              float32 `json:"totalSum"`
	WithdrawnDays         int     `json:"withdrawnDays"`
	WithdrawnDate         string  `json:"withdrawnDate"`
	SubmittedBy           string  `json:"submittedBy"`
}

type BillJoinBid struct {
	ID                    int     `json:"id" db:"id"`
	Template_code         int     `json:"template_code" db:"template_code"`
	Name                  string  `json:"name" db:"name"`
	Description           string  `json:"description" db:"description"`
	Owner                 Owner   `json:"owner" db:"owner"`
	SpecificationStimates string  `json:"specificationNStimates" db:"specifications_stimates"`
	NotIncluded           string  `json:"notIncluded" db:"not_included"`
	TotalSum              float32 `json:"totalSum" db:"totalSum"`
	WithdrawnDays         int     `json:"withdrawnDays" db:"withdrawn_days"`
	WithdrawnDate         string  `json:"withdrawnDate" db:"withdrawn_date"`
	SubmittedBy           string  `json:"submittedBy" db:"submitted_by"`
	LastEdit              string  `json:"last_edit" db:"lastEdit"`
}