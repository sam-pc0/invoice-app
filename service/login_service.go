package service

import (
	"log"

	"github.com/sam-pc0/invoice-app/model"
	"github.com/sam-pc0/invoice-app/repository"
)

type LoginService interface {
	VerifyUser() (error)
}

type DefaultLoginService struct {
	Repo repository.LoginRepository
}

func NewLoginService(repo repository.LoginRepository) DefaultLoginService {
	return DefaultLoginService{Repo: repo}
}

func (defaultService DefaultLoginService) VerifyUser(loginData model.LoginData)  (error) {
	err := defaultService.Repo.VerifyUser(loginData)
	if err != nil {
		log.Println("[LoginService Error]", err)
		return err
	}
	return nil 
}

