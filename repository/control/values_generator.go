package control

import "github.com/canxega/invoice-app/model"

func GenerateItem() model.Item {
	return model.Item{
		Amount:      0.0,
		Description: "",
	}
}

func GenerateInvoice() model.Invoice {
	return model.Invoice{
		Number:         0,
		Total:          0.0,
		DateSubmmitted: "",
	}
}

func GenerateBidProposal() model.BidProposal {
	return model.BidProposal{
		Number:                0,
		SpecificationStimates: "",
		NotIncluded:           "",
		TotalSum:              0,
		WithdrawnDays:         0,
		WithdrawnDate:         "",
	}
}
