package service

import (
	"log"

	"github.com/sam-pc0/invoice-app/model"
	"github.com/sam-pc0/invoice-app/repository"
)

type OwnerService interface {
	CreateOwner(int, model.Owner) (int, error)
	DeleteOwnerByBillId(int) (error)
	UpdateOwnerByBillId(int, model.Owner) (int, error)
}

type DefaultOwnerService struct {
	R repository.OwnerRepository
}

func NewOwnerService(r repository.OwnerRepository) DefaultOwnerService {
	return DefaultOwnerService{R: r}
}

func (s DefaultOwnerService) CreateOwner(o model.Owner) (int, error) {
	ownerId, err := s.R.CreateOwner(o)
	if err != nil {
		log.Println("[OwnerService Error]", err)
		return 0, err
	}
	return ownerId, nil
}

func (s DefaultOwnerService) DeleteOwnerbyBillId(billId int) (error) {
	err := s.R.DeleteOwnerbyBillId(billId)
	if err != nil {
		log.Println("[OwnerService Error]", err)
		return err
	}
	return nil
}

func (s DefaultOwnerService) UpdateOwnerByBillId(billId int, owner model.Owner) (int, error) {
	ownerId, err := s.R.UpdateOwnerByBillId(billId, owner) 
	if err != nil {
		log.Println("[OwnerService Error]", err)
		return 0, err
	}
	return ownerId, nil
}
