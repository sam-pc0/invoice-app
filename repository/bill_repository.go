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

func (r *BillRepository) GetAllBills() ([]model.BillRequestGet, error) {
	query := `SELECT id,name,description,template_code,lastEdit
	FROM bills`

	var b []model.BillRequestGet
	err := r.client.Select(&b, query)
	if err != nil {
		log.Println("[BillRepository Error]", err)
		return nil, err
	}

	return b, nil
}

func (r *BillRepository) CreateBill(b model.Bill) (int, error) {
	query := `INSERT INTO bills (name, description, template_code, lastEdit) VALUES (?, ?, ?, ?)`

	tx := r.client.MustBegin()
	tx.MustExec(query, b.Name, b.Description, b.TemplateCode, b.LastEdit)
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
	case 1110:
		inv := control.GenerateInvoice()
		r.UpdateBillAndCreateInvoice(b.Owner, b, inv, b.TemplateCode)
	case 1100:
		bid := control.GenerateBidProposal()
		_ = r.UpdateBillAndCreateBid(b.Owner, b, bid, b.TemplateCode)
	}

	return lastId, nil
}


func (r *BillRepository) DeleteBidById(id int) (error) {
	query := `delete bid.*, o.*, b.*
	 	from bid_proposal bid
	 	join bills b on b.id = bid.id_bill
	 	join owner o on b.owner_id = o.id
	 	and b.id =?`
	
	tx := r.client.MustBegin()
	tx.MustExec(query, id)
	if err := tx.Commit(); err != nil {
		log.Println("[BillRepository Error]", err)
		tx.Rollback()
		return err
	}

	return nil
}

func (r *BillRepository) DeleteInvoiceById(id int) (error) {
	deleteItemsQuery := `delete i, i2
        from item_invoice i
        join items i2 on i.item_id = i2.id
        and i.invoice_item =?`
	
	deleteBillQuery := `delete i.*, o.*, b.*
	 	from invoices i
	 	join bills b on b.id = i.id_bill
	 	join owner o on b.owner_id = o.id
	 	and b.id = ?`
	
	tx := r.client.MustBegin()
	tx.MustExec(deleteItemsQuery, id)
	if err := tx.Commit(); err != nil {
		log.Println("[BillRepository Error]", err)
		tx.Rollback()
		return err
	}

	tx2 := r.client.MustBegin()
	tx2.MustExec(deleteBillQuery, id)
	if err := tx2.Commit(); err != nil {
		log.Println("[BillRepository Error]", err)
		tx.Rollback()
		return err
	}
	

	return nil
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

func (r *BillRepository) GetBillContentByID(id int) (model.BillJionBid, error) {
	dbBid := NewBidProposalRepository(r.client)
	b, err := dbBid.GetBidAndBillByID(id)
	if err != nil {
		log.Println("[BillRepository Error]", err)
		return model.BillJionBid{}, err
	}

	return b, nil
}



func (r *BillRepository) UpdateBill(b model.Bill, id int) error {
	log.Print("is here")
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
	for n, i := range it {
		if n == 0 {
			err = dbItem.UpdateItem(i, i.ID)
		} else {
			x, _ := dbItem.SaveItem(i)
			err = dbItem.SaveItemInvoice(x, in.ID)
		}
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

func (r *BillRepository) UpdateBillAndCreateInvoice(owner model.Owner, bil model.Bill, invoice model.Invoice, code int) error {
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

	invoiceService := NewInvoiceRepository(r.client)
	invoiceId, err := invoiceService.SaveInvoice(invoice, bil.ID)
	if err != nil {
		log.Println("[BillRepository Error]", err)
		return err
	}

	log.Print(invoiceId)

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
