package service

import (
	"github.com/canxega/invoice-app/model"
	"github.com/canxega/invoice-app/repository"
)

type BillServie interface {
	GetAllBillS() ([]model.BillRequestGet, error)
	SaveBasicBill(model.Bill) (int, error)
	GetBillById(int) (model.Bill, error)
	SaveBillBid(model.Owner, model.Bill, model.BidProposal, int) error
	SaveBillInvoice(model.Owner, model.Bill, model.Invoice, []model.Item, int) error
	GetBidBill(int) (model.BillJionBid, error)
	GetInvoiceBill(int) (model.BillJoinInvoice, error)
	VerifyCode(int) (int, error)
	UpdateBillBid(model.Owner, model.Bill, model.BidProposal) error
	UpdateBillInvoice(model.Owner, model.Bill, model.Invoice, []model.Item) error
}

type DefaultBillService struct {
	R repository.BillRepository
}

func NewBillService(r repository.BillRepository) DefaultBillService {
	return DefaultBillService{R: r}
}

func (s DefaultBillService) GetAllBillS() ([]model.BillRequestGet, error) {
	return s.R.GetAllBills()
}

func (s DefaultBillService) SaveBasicBill(b model.Bill) (int, error) {
	return s.R.InsertContent(b)
}

func (s DefaultBillService) GetBillById(id int) (model.Bill, error) {
	return s.R.GetBillByID(id)
}

func (s DefaultBillService) SaveBillBid(o model.Owner, bill model.Bill, bid model.BidProposal, code int) error {
	return s.R.UpdateBillAndCreateBid(o, bill, bid, code)
}

func (s DefaultBillService) SaveBillInvoice(o model.Owner, b model.Bill, inv model.Invoice, it []model.Item, code int) error {
	return s.R.UpdateBillAndCreateInvoice(o, b, inv, code)
}

func (s DefaultBillService) GetBidBill(id int) (model.BillJionBid, error) {
	return s.R.GetBillContentByID(id)
}

func (s DefaultBillService) GetInvoiceBill(id int) (model.BillJoinInvoice, error) {
	return s.R.GetBillInvoiceContentByID(id)
}

func (s DefaultBillService) VerifyCode(id int) (int, error) {
	return s.R.VerifyCode(id)
}

func (s DefaultBillService) UpdateBillBid(o model.Owner, b model.Bill, bid model.BidProposal) error {
	return s.R.UpdateContentBid(o, b, bid)
}

func (s DefaultBillService) UpdateBillInvoice(o model.Owner, b model.Bill, in model.Invoice, it []model.Item) error {
	return s.R.UpdateContentInvoice(o, b, in, it)
}
