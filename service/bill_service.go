package service

import (
	"log"

	"github.com/canxega/invoice-app/model"
	"github.com/canxega/invoice-app/repository"
)

type BillService interface {
	GetAllBillS() ([]model.BillRequestGet, error)
	CreateBill(int, model.Bill) (int, error)
	GetBillById(int) (model.Bill, error)
	GetTemplateCode(int) (int, error)
	DeleteById(int) (error)

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
	bills, err := s.R.GetAllBills()
	if err != nil {
		log.Println("[BillService Error]", err)
		return nil, err
	}
	return bills, nil
}

func (s DefaultBillService) CreateBill(ownerId int, b model.Bill) (int, error) {
	billId, err := s.R.CreateBill(ownerId, b)
	if err != nil {
		log.Println("[BillService Error]", err)
		return 0, err
	}
	return billId, err
}

func (s DefaultBillService) DeleteById(billId int) (error) {
	err := s.R.DeleteById(billId)
	if err != nil {
		log.Println("[BillService Error]", err)
		return  err
	}
	return nil
}

func (s DefaultBillService) GetTemplateCode(id int) (int, error) {
	return s.R.GetBillTemplateCode(id)
}

func (s DefaultBillService) GetBillById(id int) (model.Bill, error) {
	return s.R.GetBillByID(id)
}


func (s DefaultBillService) UpdateBillBid(o model.Owner, b model.Bill, bid model.BidProposal) error {
	return s.R.UpdateContentBid(o, b, bid)
}

func (s DefaultBillService) UpdateBillInvoice(o model.Owner, b model.Bill, in model.Invoice, it []model.Item) error {
	return s.R.UpdateContentInvoice(o, b, in, it)
}
