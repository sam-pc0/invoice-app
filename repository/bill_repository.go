package repository

import (
	"log"

	"github.com/canxega/invoice-app/model"
	"github.com/canxega/invoice-app/repository/control"
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
	tx.MustExec(query, b.Name, b.Description, b.TemplateCode)
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

	b.ID = lastId
	switch b.TemplateCode {
	case 1100:
		bid := control.GenerateBidProposal()
		_ = r.UpdateBillAndCreateBid(b.Owner, b, bid, b.TemplateCode)
	case 1110:
		inv := control.GenerateInvoice()
		items := []model.Item{}
		items = append(items, control.GenerateItem())
		_ = r.UpdateBillAndCreateInvoice(b.Owner, b, inv, items, b.TemplateCode)
	}

	return lastId, nil
}

func (r *BillRepository) GetBillByID(id int) (model.Bill, error) {
	query := `SELECT id, 
	name,
	description,
	template_code
	FROM bills WHERE id = ?`

	var b model.Bill
	err := r.client.Get(&b, query, id)
	if err != nil {
		log.Println("[BillRepository Error]", err)
		return model.Bill{}, err
	}

	return b, nil
}

func (r *BillRepository) GetAllBills() ([]model.BillRequestGet, error) {
	query := `SELECT id,name,description,template_code
	FROM bills`

	var b []model.BillRequestGet
	err := r.client.Select(&b, query)
	if err != nil {
		log.Println("[BillRepository Error]", err)
		return nil, err
	}

	return b, nil
}

func (r *BillRepository) UpdateBill(b model.Bill, id int) error {
	query := `UPDATE bills SET
	name = ?,
	description = ?,
	lastEdit = ?
	WHERE id=?`

	tx := r.client.MustBegin()
	tx.MustExec(query, b.Name, b.Description, b.LastEdit, id)
	if err := tx.Commit(); err != nil {
		log.Println("[BillRepository Error]", err)
		tx.Rollback()
		return err
	}

	return nil
}

func (r *BillRepository) GetBillContentByID(id int) (model.BillJionBid, error) {
	dbBid := NewBidProposalRepository(r.client)
	b, err := dbBid.GetBidAndBillByID(id)
	if err != nil {
		log.Println("[BillRepository Error]", err)
		return model.BillJionBid{}, err
	}

	return b, nil
}

func (r *BillRepository) GetBillInvoiceContentByID(id int) (model.BillJoinInvoice, error) {
	dbInv := NewInvoiceRepository(r.client)
	b, err := dbInv.GetInvoiceAndBillByID(id)
	if err != nil {
		log.Println("[BillRepository Error]", err)
		return model.BillJoinInvoice{}, err
	}

	return b, nil
}

func (r *BillRepository) VerifyCode(id int) (int, error) {
	query := `SELECT 
	template_code
	FROM bills WHERE id = ?`

	var code int
	err := r.client.Get(&code, query, id)
	if err != nil {
		log.Println("[BillRepository Error]", err)
		return 0, err
	}

	return code, nil
}

func (r *BillRepository) UpdateContentBid(o model.Owner, b model.Bill, bid model.BidProposal) error {
	dbOwner := NewOwnerRepository(r.client)
	err := dbOwner.UpdateOwner(o, o.ID)
	if err != nil {
		log.Println("[BillRepository Error]", err)
		return err
	}

	dbBill := NewBillRepository(r.client)
	err = dbBill.UpdateBill(b, b.ID)
	if err != nil {
		log.Println("[BillRepository Error]", err)
		return err
	}

	dbBid := NewBidProposalRepository(r.client)
	err = dbBid.UpdateBid(bid, bid.ID)
	if err != nil {
		log.Println("[BillRepository Error]", err)
		return err
	}

	return nil
}

func (r *BillRepository) UpdateContentInvoice(o model.Owner, b model.Bill, in model.Invoice, it []model.Item) error {
	dbOwner := NewOwnerRepository(r.client)
	err := dbOwner.UpdateOwner(o, o.ID)
	if err != nil {
		log.Println("[BillRepository Error]", err)
		return err
	}

	dbBill := NewBillRepository(r.client)
	err = dbBill.UpdateBill(b, b.ID)
	if err != nil {
		log.Println("[BillRepository Error]", err)
		return err
	}

	dbInv := NewInvoiceRepository(r.client)
	err = dbInv.UpdateInvoice(in, in.ID)
	if err != nil {
		log.Println("[BillRepository Error]", err)
		return err
	}

	dbItem := NewItemRepository(r.client)
	for _, i := range it {
		err = dbItem.UpdateItem(i, i.ID)
		if err != nil {
			log.Println("[BillRepository Error]", err)
			return err
		}
	}

	return nil
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
	save.SaveBidProposal(bid, bill.ID)

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
	idInvoice, err := save.SaveInvoice(invoice, bil.ID)
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
