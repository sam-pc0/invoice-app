package model

type BidProposal struct {
	ID                    int     `json:"id"`
	SpecificationStimates string  `json:"specificationstimates"`
	NotIncluded           string  `json:"notIncluded"`
	TotalSum              float32 `json:"totalsum"`
	WithdrawnDays         int     `json:"withdrawnDays"`
	WithdrawnDate         string  `json:"withdrawnDate"`
}
