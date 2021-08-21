package service

import (
	"github.com/canxega/invoice-app/model"
	"github.com/canxega/invoice-app/repository"
)

type OwnerService interface {
	SaveOwner(model.Owner) (int, error)
	GetOwenerByID(int) (model.Owner, error)
}

type DefaultOwnerService struct {
	R repository.OwnerRepository
}

func NewOwnerService(r repository.OwnerRepository) DefaultOwnerService {
	return DefaultOwnerService{R: r}
}

func (s DefaultOwnerService) SaveOwner(o model.Owner) (int, error) {
	return s.R.SaveOwner(o)
}

func (s DefaultOwnerService) GetOwenerByID(id int) (model.Owner, error) {
	return s.R.GetOwnerByID(id)
}
