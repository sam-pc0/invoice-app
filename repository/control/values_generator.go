package control

import "github.com/sam-pc0/invoice-app/model"

func GenerateItem() model.Item {
	return model.Item{
		Item:        "",
		Amount:      0.0,
		Description: "",
	}
}

func GenerateInvoice() model.Invoice {
	return model.Invoice{
		DateSubmmitted: "",
	}
}

func GenerateBidProposal() model.BidProposal {
	return model.BidProposal{
		SpecificationStimates: "",
		NotIncluded:           "",
		WithdrawnDays:         0,
		WithdrawnDate:         "",
	}
}
