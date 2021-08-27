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

func (r *BillRepository) CreateBill(ownerId int, b model.Bill) (int, error) {
	query := `INSERT INTO bills (name, description, template_code, lastEdit, owner_id) VALUES (?, ?, ?, ?, ?)`

	tx := r.client.MustBegin()
	tx.MustExec(query, b.Name, b.Description, b.TemplateCode, b.LastEdit, ownerId)
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

	return lastId, nil
}

func (r *BillRepository) GetBillTemplateCode(id int) (int, error) {
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

func (r *BillRepository) DeleteById(id int) (error) {
	query := `delete b from bills b where b.id = ?`

	tx := r.client.MustBegin()
	tx.MustExec(query, id)
	if err := tx.Commit(); err != nil {
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

	// dbItem := NewItemRepository(r.client)
	// for n, i := range it {
	// 	if n == 0 {
	// 		err = dbItem.UpdateItem(i, i.ID)
	// 	} else {
	// 		x, _ := dbItem.SaveItem(i)
	// 		err = dbItem.SaveItemInvoice(x, in.ID)
	// 	}
	// 	if err != nil {
	// 		log.Println("[BillRepository Error]", err)
	// 		return err
	// 	}
	// }

	return nil
}
