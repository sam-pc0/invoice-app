package model

type BidProposal struct {
	ID                    int      `json:"id"`
	TemplateCode          Template `json:"template" db:"template_code"`
	Name                  string   `json:"name"`
	Description           string   `json:"description"`
	Number                int      `json:"number"`
	SpecificationStimates string   `json:"specificationstimates"`
	NotIncluded           string   `json:"not_included"`
	TotalSum              float32  `json:"totalsum"`
	WithdrawnDays         string   `json:"withdrawndays"`
	WithdrawnDate         string   `json:"withdrawndate"`
}
