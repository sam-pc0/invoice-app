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

func (r *OwnerRepository) CreateOwner(o model.Owner) (int, error) {
	query := `INSERT INTO owner(name, location, phone, altPhone, projectNameAddress, email, address)
	VALUES (?, ?, ?, ?, ?, ?, ?)`

	tx := r.client.MustBegin()
	tx.MustExec(query, o.Name, o.Location, o.Phone, o.AltPhone, o.ProjectNameAddress, o.Email, o.Address)
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

func (r *OwnerRepository) DeleteOwnerbyBillId(billId int) (error) {
	query := `DELETE o FROM owner o JOIN bills b ON b.owner_id = o.id WHERE b.id = ?`
	tx := r.client.MustBegin()
	tx.MustExec(query, billId)
	if err := tx.Commit(); err != nil {
		log.Println("[OwnerRepository Error]", err)
		tx.Rollback()
		return err
	}
	return nil
}

func (r *OwnerRepository) UpdateOwnerByBillId(billId int, o model.Owner) (int, error) {
	query := `
	UPDATE owner 
	JOIN bills b ON b.owner_id=owner.id 
	SET
	owner.name = ?,
	owner.location = ?,
	owner.phone = ?,
	owner.altPhone = ?,
	owner.projectNameAddress = ?,
	owner.email = ?,
	owner.address = ?
	WHERE b.id =?`

	tx := r.client.MustBegin()
	tx.MustExec(query, o.Name, o.Location, o.Phone, o.AltPhone, o.ProjectNameAddress, o.Email, o.Address, billId)
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
