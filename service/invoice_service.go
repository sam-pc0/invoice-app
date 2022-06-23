package service

import (
	"log"

	"github.com/sam-pc0/invoice-app/model"
	"github.com/sam-pc0/invoice-app/repository"
	"github.com/sam-pc0/invoice-app/repository/control"
)

type InvoiceService interface {
	CreateInvoice(int) (int, error)
	GetFullInvoiceByBillId(int) (model.BillJoinInvoice)
	DeleteByBillId(int) (error)
	UpdateInvoice(int, model.Invoice, []model.Item)
}

type DefaultInvoiceService struct {
	R repository.InvoiceRepository
}

func NewInvoiceService(r repository.InvoiceRepository) DefaultInvoiceService{
	return DefaultInvoiceService{R: r}
}

func (s DefaultInvoiceService) CreateInvoice(billId int) (int, error) {
	invoice := control.GenerateInvoice()
	invoiceId, err := s.R.CreateInvoice(billId, invoice)
	if err != nil {
		log.Println("[InvoiceService Error]", err)
		return 0, err
	}

	return invoiceId, nil
}

func (s DefaultInvoiceService) GetFullInvoiceByBillId(billId int) (model.BillJoinInvoice, error) {
	billInvoice, err := s.R.GetInvoiceByBillId(billId)
	if err != nil {
		log.Println("[InvoiceService Error]", err)
		return model.BillJoinInvoice{}, err
	}
	items, _ := s.GetItemsByBillId(billId)
	billInvoice.Items = items

	return billInvoice, nil
}

func (s DefaultInvoiceService) DeleteByBillId(billId int) (error) {
	err := s.DeleteAllItems(billId)
	if err != nil {
		log.Println("[InvoiceService Error]", err)
		return err
	}
	
	err = s.R.DeleteByBillId(billId)
	if err != nil {
		log.Println("[InvoiceService Error]", err)
		return err
	}
	return nil 
}

func (s DefaultInvoiceService) UpdateInvoice(billId int, i model.Invoice, items []model.Item) (error) {
	invoiceId, err := s.R.UpdateInvoice(billId, i)
	if err != nil {
		log.Println("[InvoiceService Error]", err)
		return err
	}

    err = s.DeleteAllItems(billId)
	if err != nil {
		log.Println("[InvoiceService Error]", err)
		return err
	}

    err = s.CreateItems(invoiceId, items)
	if err != nil {
		log.Println("[InvoiceService Error]", err)
		return err
	}

	return nil 
}

func (s DefaultInvoiceService) CreateItems(billId int, items []model.Item) (error) {
	for _, i := range items {
		err := s.R.CreateItem(billId, i)
		if err != nil {
			log.Println("[BillRepository Error]", err)
			return err
		}
	}


	return nil
}

func (s DefaultInvoiceService) GetItemsByBillId(billId int) ([]model.Item, error) {
	items, err := s.R.GetItemsByBillId(billId)
	if err != nil {
		log.Println("[InvoiceService Error]", err)
		return nil, err
	}
	return items, nil
}

func (s DefaultInvoiceService) DeleteAllItems(billId int) (error) {
	err := s.R.DeleteAllItems(billId)
	if err != nil {
		log.Println("[InvoiceService Error]", err)
		return err
	}
	return nil 
}

