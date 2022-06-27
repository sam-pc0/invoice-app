package repository

import (
	"github.com/sam-pc0/invoice-app/model"
	"github.com/jmoiron/sqlx"
)

type LoginRepository struct {
	client *sqlx.DB
}

func NewLoginRepository(db *sqlx.DB) LoginRepository {
	return LoginRepository{db}
}

func (repo *LoginRepository) VerifyUser(loginData model.LoginData) (error) {
	var name string
	err := repo.client.Get(&name, "SELECT name FROM users WHERE name=? AND password=?", loginData.Name, loginData.Password)
	if err != nil	{
		return err;
	}
	return nil
}

