package repository

import (
	"log"

	"github.com/canxega/invoice-app/model"
	"github.com/jmoiron/sqlx"
)

type InvoiceRepository struct {
	client *sqlx.DB
}

func NewInvoiceRepository(db *sqlx.DB) InvoiceRepository {
	return InvoiceRepository{db}
}

func (r *InvoiceRepository) SaveInvoice(invoice model.Invoice, id int) (int, error) {
	query := `INSERT INTO invoices (number_inv, total, dateSubmmitted, id_bill)
	VALUES (?,?,?,?)`

	tx := r.client.MustBegin()
	tx.MustExec(query, invoice.Number, invoice.Total, invoice.DateSubmmitted, id)
	if err := tx.Commit(); err != nil {
		log.Println("[InvoiceRepository Error]", err)
		tx.Rollback()
		return 0, err
	}

	var lastId int
	err := r.client.Get(&lastId, "SELECT LAST_INSERT_ID()")
	if err != nil {
		log.Println("[InvoiceRepository Error]", err)
		return 0, err
	}

	return lastId, nil

}
