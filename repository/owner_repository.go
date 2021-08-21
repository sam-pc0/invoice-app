package repository

import "github.com/jmoiron/sqlx"

type OwnerRepository struct {
	client *sqlx.DB
}

func NewOwnerRepository(db *sqlx.DB) OwnerRepository {
	return OwnerRepository{db}
}
