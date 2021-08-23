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

func (r *InvoiceRepository) GetInvoiceAndBillByID(id int) (model.BillJoinInvoice, error) {
	query := `SELECT bills.id,bills.template_code,
	bills.name, bills.description, bills.lastEdit,
	owner.id, owner.name, owner.location, owner.phone, owner.altPhone, owner.projectNameAddress, owner.email,
	invoices.id, invoices.number_inv, invoices.total, invoices.dateSubmmitted,
	item_invoice.id, items.id, items.description, items.amount 
	FROM bills  INNER JOIN owner  
	ON bills.owner_id = owner.id
	INNER JOIN invoices  ON invoices.id_bill = bills.id 
	INNER JOIN item_invoice  ON item_invoice.invoice_item = invoices.id 
	INNER JOIN items  ON items.id  = item_invoice.item_id 
	WHERE bills.id =?`

	var b model.BillJoinInvoice
	err := r.client.Get(&b, query, id)
	if err != nil {
		log.Println("[BidProposalRepository]", err)
		return model.BillJoinInvoice{}, err
	}

	return b, nil
}
