package repository

import (
	"log"

	"github.com/canxega/invoice-app/model"
	"github.com/jmoiron/sqlx"
)

type BillRepository struct {
	client *sqlx.DB
}

func NewBillRepository(db *sqlx.DB) BillRepository {
	return BillRepository{db}
}

func (r *BillRepository) InsertContent(b model.Bill) (int, error) {
	query := `INSERT INTO bills (name, description, template_code) VALUES (?, ?, ?)`

	tx := r.client.MustBegin()
	tx.MustExec(query, b.Name, b.Description, b.Template.TemplateCode)
	if err := tx.Commit(); err != nil {
		log.Println("[BillRepository Error]", err)
		tx.Rollback()
		return 0, err
	}

	var lastId int
	err := r.client.Get(&lastId, "SELECT LAST_INSERT_ID()")
	if err != nil {
		log.Println("[BillRepository Error]", err)
		return 0, err
	}

	return lastId, nil
}

func (r *BillRepository) GetBillByID(id int) (model.Bill, error) {
	query := `SELECT id, 
	name, template_code "template_code.id",
	description
	FROM bills WHERE id = ?`

	var b model.Bill
	err := r.client.Get(&b, query, id)
	if err != nil {
		log.Println("[BillRepository Error]", err)
		return model.Bill{}, err
	}

	return b, nil
}
