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

func (r *InvoiceRepository) CreateInvoice(billId int, invoice model.Invoice) (int, error) {
	query := `INSERT INTO invoices (dateSubmmitted, id_bill)
	VALUES (?,?)`

	tx := r.client.MustBegin()
	tx.MustExec(query, invoice.DateSubmmitted, billId)
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
	return lastId, err 
}

func (r *InvoiceRepository) DeleteByBillId(billId int) (error) {
	query := `delete i
	 	from invoices i
	 	join bills b on b.id = i.id_bill
	 	and b.id = ?`
	
	tx := r.client.MustBegin()
	tx.MustExec(query, billId)
	if err := tx.Commit(); err != nil {
		log.Println("[BillRepository Error]", err)
		tx.Rollback()
		return err
	}
	return nil
}

func (r *InvoiceRepository) GetInvoiceByBillId(id int) (model.BillJoinInvoice, error) {
	query := `
  SELECT
		bills.id,
		bills.template_code,
		bills.name, 
		bills.description, 
		bills.lastEdit,
		bills.total,	
		bills.sub_total,
		bills.tax_rate,
		bills.tax,
		owner.name "owner.name",
		owner.location "owner.location", 
		owner.phone "owner.phone", 
		owner.altPhone "owner.altPhone",
		owner.projectNameAddress "owner.projectNameAddress",
		owner.address "owner.address",
		owner.email "owner.email",
		invoices.dateSubmmitted
	FROM bills
	JOIN owner ON bills.owner_id = owner.id
	JOIN invoices ON invoices.id_bill = bills.id 
	AND bills.id =?`

	var b model.BillJoinInvoice
	err := r.client.Get(&b, query, id)
	if err != nil {
		return model.BillJoinInvoice{}, err
	}

	return b, nil
}

func (r *InvoiceRepository) UpdateInvoice(billId int, i model.Invoice) (int, error) {
	query := `
	UPDATE invoices 
	JOIN bills b ON b.id = invoices.id_bill
	SET
	invoices.dateSubmmitted = ?
	WHERE b.id = ?`

	log.Print(i.DateSubmmitted)
	tx := r.client.MustBegin()
	tx.MustExec(query, i.DateSubmmitted, billId)
	if err := tx.Commit(); err != nil {
		log.Println("[InvoiceRepository Error]", err)
		tx.Rollback()
		return 0, err
	}

	var lastId int
	err := r.client.Get(&lastId, "SELECT i.id FROM invoices i JOIN bills b ON b.id = i.id_bill WHERE b.id = ?", billId)
	if err != nil {
		log.Println("[InvoiceRepository Error]", err)
		return 0, err
	}

	return lastId, nil
}

func (r *InvoiceRepository) GetItemsByBillId(id int) ([]model.Item, error) {
	query := `
		SELECT
		items.item,
		items.amount,
		items.description
			FROM
		items
		JOIN item_invoice ii ON items.id = ii.item_id
		JOIN invoices i ON i.id = ii.invoice_item
		JOIN bills b ON b.id = i.id_bill
			AND b.id = ?`

   i := []model.Item{}
   err := r.client.Select(&i, query, id)
   if err != nil {
	   log.Println("[InvoiceRepository]", err)
	   return nil, err
   }
   return i, nil
}

func (r *InvoiceRepository) DeleteAllItems(billId int) (error) {
	deleteItemsQuery := `
	DELETE i, i2 
	FROM item_invoice i
	JOIN items i2 ON i.item_id = i2.id
	JOIN invoices i3 ON i3.id = i.invoice_item
	JOIN bills b ON b.id = i3.id_bill
	AND b.id = ?`

	tx := r.client.MustBegin()
	tx.MustExec(deleteItemsQuery, billId)
	if err := tx.Commit(); err != nil {
		log.Println("[ItemsRepository Error]", err)
		tx.Rollback()
		return err
	}
	return nil
}

func (r *InvoiceRepository) CreateItem(invoiceId int, item model.Item) (error) {
	query := `INSERT INTO items (description, amount, item)
	VALUES (?,?,?)`

	tx := r.client.MustBegin()
	tx.MustExec(query, item.Description, item.Amount, item.Item)
	if err := tx.Commit(); err != nil {
		log.Println("[ItemsRepository Error]", err)
		tx.Rollback()
		return err
	}

	var itemId int
	err := r.client.Get(&itemId, "SELECT LAST_INSERT_ID()")
	if err != nil {
		log.Println("[InvoiceRepository Error]", err)
		return err
	}

	return r.ItemRelation(invoiceId, itemId)
}

func (r *InvoiceRepository) ItemRelation(invoiceId int, itemId int) (error) {
	query := `INSERT INTO item_invoice (item_id, invoice_item)
	VALUES (?,?)`

	tx := r.client.MustBegin()
	tx.MustExec(query, itemId, invoiceId)
	if err := tx.Commit(); err != nil {
		log.Println("[ItemsRepository Error]", err)
		tx.Rollback()
		return err
	}
	return nil
}
