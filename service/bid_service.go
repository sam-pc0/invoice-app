package service

import (
	"log"

	"github.com/canxega/invoice-app/model"
	"github.com/canxega/invoice-app/repository"
	"github.com/canxega/invoice-app/repository/control"
)

type BidService interface {
	CreateBid(int) (int, error)
	GetFullInvoiceByBillId(int) (model.BillJoinBid)
	DeleteByBillId(int) (error)
	UpdateBid(int, model.BidProposal)
}

type DefaultBidService struct {
	R repository.BidProposalRepository
}

func NewBidService(r repository.BidProposalRepository) DefaultBidService {
	return DefaultBidService{R: r}
}

func (s DefaultBidService) CreateBid(billId int) (int, error) {
	bidProposal := control.GenerateBidProposal()
	billId, err := s.R.CreateBid(billId, bidProposal )
	if err != nil {
		log.Println("[BidService Error]", err)
		return 0, err
	}
	return billId, err

}

func (s DefaultBidService) GetFullBidByBillId(billId int) (model.BillJoinBid, error) {
	billBid, err := s.R.GetFullBidByBillId(billId)
	if err != nil {
		log.Println("[BidService Error]", err)
		return model.BillJoinBid{}, err
	}
	return billBid, nil
}

func (s DefaultBidService) DeleteByBillId(billid int) (error) {
	err := s.R.DeleteByBillId(billid) 
	if err != nil {
		log.Println("[InvoiceService Error]", err)
		return err
	}
	return nil
}

func (s DefaultBidService) UpdateBid(billid int, b model.BidProposal) (error) {
	err := s.R.UpdateBid(billid, b) 
	if err != nil {
		log.Println("[InvoiceService Error]", err)
		return err
	}
	return nil
}