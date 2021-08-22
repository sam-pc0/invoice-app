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
	name, template_code "template_code.templatecode",
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

func (r *BillRepository) UpdateBillAndCreateBid(owner model.Owner, bill model.Bill, bid model.BidProposal, code int) error {
	id, err := saveOwnerBill(r.client, owner)
	if err != nil {
		log.Println("[BillRepository Error]", err)
		return err
	}

	err = updateBillOperation(r.client, code, id, bill.ID)
	if err != nil {
		log.Println("[BillRepository Error]", err)
		return err
	}

	save := NewBidProposalRepository(r.client)
	save.SaveBidProposal(bid)

	return nil
}

func (r *BillRepository) UpdateBillAndCreateInvoice(owner model.Owner, bil model.Bill, invoice model.Invoice, items []model.Item, code int) error {
	id, err := saveOwnerBill(r.client, owner)
	if err != nil {
		log.Println("[BillRepository Error]", err)
		return err
	}

	err = updateBillOperation(r.client, code, id, bil.ID)
	if err != nil {
		log.Println("[BillRepository Error]", err)
		return err
	}

	save := NewInvoiceRepository(r.client)
	idInvoice, err := save.SaveInvoice(invoice)
	if err != nil {
		log.Println("[BillRepository Error]", err)
		return err
	}

	dbItem := NewItemRepository(r.client)
	for _, i := range items {
		idItem, err := dbItem.SaveItem(i)
		if err != nil {
			log.Println("[BillRepository Error]", err)
			return err
		}
		if err := dbItem.SaveItemInvoice(idItem, idInvoice); err != nil {
			log.Println("[BillRepository Error]", err)
			return err
		}
	}

	return nil
}

func saveOwnerBill(c *sqlx.DB, owner model.Owner) (int, error) {
	db := NewOwnerRepository(c)
	return db.SaveOwner(owner)
}

func updateBillOperation(c *sqlx.DB, code, idO, idB int) error {
	query := `UPDATE bills SET 
	template_code = ?,
	owner_id=?
	WHERE id=?`

	tx := c.MustBegin()
	tx.MustExec(query, code, idO, idB)

	if err := tx.Commit(); err != nil {
		log.Println("[BillRepository Error]", err)
		tx.Rollback()
		return err
	}

	return nil
}
