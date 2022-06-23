package service

import (
	"log"

	"github.com/sam-pc0/invoice-app/model"
	"github.com/sam-pc0/invoice-app/repository"
)

type BillService interface {
	GetAllBillS() ([]model.BillRequestGet, error)
	CreateBill(int, model.Bill) (int, error)
	GetTemplateCode(int) (int, error)
	DeleteById(int) (error)
	UpdateBill(int, model.Bill) (error)
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

func (s DefaultBillService) UpdateBill(billId int,  bill model.Bill) error {
	return s.R.UpdateBill(billId, bill)
}
