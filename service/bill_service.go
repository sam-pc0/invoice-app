package service

import (
	"github.com/canxega/invoice-app/model"
	"github.com/canxega/invoice-app/repository"
)

type BillServie interface {
	SaveBasicBill(model.Bill) (int, error)
	GetBillById(int) (model.Bill, error)
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

func (s DefaultBillService) GetBillById(id int) (model.Bill, error) {
	return s.R.GetBillByID(id)
}
