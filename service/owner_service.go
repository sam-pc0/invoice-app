package service

import (
	"log"

	"github.com/canxega/invoice-app/model"
	"github.com/canxega/invoice-app/repository"
)

type OwnerService interface {
	CreateOwner(int, model.Owner) (int, error)
	DeleteOwnerbyBillId(int) (error)
	GetOwenerByID(int) (model.Owner, error)
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

func (s DefaultOwnerService) GetOwenerByID(id int) (model.Owner, error) {
	owner, err := s.R.GetOwnerByID(id)
	if err != nil {
		log.Println("[OwnerService Error]", err)
		return model.Owner{}, err
	}
	return owner, nil

}
