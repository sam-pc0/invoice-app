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
	query := `INSERT INTO invoices (total, dateSubmmitted, id_bill)
	VALUES (?,?,?)`

	tx := r.client.MustBegin()
	tx.MustExec(query, invoice.Total, invoice.DateSubmmitted, id)
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
	query := `SELECT bills.id "id_bill",bills.template_code,
	bills.name , bills.description, bills.lastEdit,
	owner.id "owner.id", owner.name "owner.name", owner.location "owner.location", owner.phone "owner.phone", 
	owner.altPhone "owner.altPhone", owner.projectNameAddress "owner.projectNameAddress", owner.address "owner.address", owner.email "owner.email",
	invoices.id "invoice_id", invoices.total, invoices.dateSubmmitted "date_submmitted"
	FROM bills  INNER JOIN owner  
	ON bills.owner_id = owner.id
	INNER JOIN invoices  ON invoices.id_bill = bills.id 
	INNER JOIN item_invoice  ON item_invoice.invoice_item = invoices.id 
	INNER JOIN items  ON items.id  = item_invoice.item_id 
	WHERE bills.id =?`

	var b model.BillJoinInvoice
	err := r.client.Get(&b, query, id)
	if err != nil {
		log.Println("[InvoiceRepository]", err)
		return model.BillJoinInvoice{}, err
	}

	i, _ := r.GetItemsByll(id)
	b.Item = i

	return b, nil
}

func (r *InvoiceRepository) GetItemsByll(id int) ([]model.Item, error) {
	query := `SELECT
	items.id "id", items.description "description", items.amount "amount", items.item
	FROM items INNER JOIN item_invoice ON item_invoice.item_id = items.id
	INNER JOIN invoices ON item_invoice.invoice_item = invoices.id
	INNER JOIN bills ON bills.id = invoices.id_bill
	WHERE bills.id = ?`
	i := []model.Item{}
	err := r.client.Select(&i, query, id)
	if err != nil {
		log.Println("[InvoiceRepository]", err)
		return nil, err
	}

	return i, nil
}

func (r *InvoiceRepository) UpdateInvoice(i model.Invoice, id int) error {
	query := `UPDATE invoices SET
	total = ?,
	dateSubmmitted = ?
	WHERE id = ?`

	tx := r.client.MustBegin()
	tx.MustExec(query, i.Total, i.DateSubmmitted, id)
	if err := tx.Commit(); err != nil {
		log.Println("[InvoiceRepository Error]", err)
		tx.Rollback()
		return err
	}

	return nil
}
