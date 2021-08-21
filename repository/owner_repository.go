package repository

import (
	"log"

	"github.com/canxega/invoice-app/model"
	"github.com/jmoiron/sqlx"
)

type OwnerRepository struct {
	client *sqlx.DB
}

func NewOwnerRepository(db *sqlx.DB) OwnerRepository {
	return OwnerRepository{db}
}

func (r *OwnerRepository) SaveOwner(o model.Owner) (int, error) {
	query := `INSERT INTO owner(name, location, phone, altPhone, projectNameAddress, email)
	VALUES (?, ?, ?, ?, ?, ?)`

	tx := r.client.MustBegin()
	tx.MustExec(query, o.Name, o.Location, o.Phone, o.AltPhone, o.ProjectNameAddress, o.Email)
	if err := tx.Commit(); err != nil {
		log.Println("[OwnerRepository Error]", err)
		tx.Rollback()
		return 0, err
	}

	var lastId int
	err := r.client.Get(&lastId, "SELECT LAST_INSERT_ID()")
	if err != nil {
		log.Println("[OwnerRepository Error]", err)
		return 0, err
	}

	return lastId, nil
}

func (r *OwnerRepository) GetOwnerByID(id int) (model.Owner, error) {
	query := `SELECT * FROM owner WHERE id=?`

	var o model.Owner
	err := r.client.Get(&o, query, id)
	if err != nil {
		log.Println("[BillRepository Error]", err)
		return model.Owner{}, err
	}

	return o, nil
}
