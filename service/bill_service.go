package service

import (
	"github.com/canxega/invoice-app/model"
	"github.com/canxega/invoice-app/repository"
)

type BillServie interface {
	SaveBasicBill(model.Bill) (int, error)
	GetAllBillS() ([]model.BillRequestGet, error)
	GetBillById(int) (model.Bill, error)
	SaveBillBid(model.Owner, model.Bill, model.BidProposal, int) error
	SaveBillInvoice(model.Owner, model.Bill, model.Invoice, []model.Item, int) error
}

type DefaultBillService struct {
	R repository.BillRepository
}

func NewBillService(r repository.BillRepository) DefaultBillService {
	return DefaultBillService{R: r}
}

func (s DefaultBillService) SaveBasicBill(b model.Bill) (int, error) {
	return s.R.InsertContent(b)
}

func (s DefaultBillService) GetAllBillS() ([]model.BillRequestGet, error) {
	return s.R.GetAllBills()
}

func (s DefaultBillService) GetBillById(id int) (model.Bill, error) {
	return s.R.GetBillByID(id)
}

func (s DefaultBillService) SaveBillBid(o model.Owner, bill model.Bill, bid model.BidProposal, code int) error {
	return s.R.UpdateBillAndCreateBid(o, bill, bid, code)
}

func (s DefaultBillService) SaveBillInvoice(o model.Owner, b model.Bill, inv model.Invoice, it []model.Item, code int) error {
	return s.R.UpdateBillAndCreateInvoice(o, b, inv, it, code)
}
