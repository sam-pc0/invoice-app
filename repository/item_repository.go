package repository

import (
	"log"

	"github.com/canxega/invoice-app/model"
	"github.com/jmoiron/sqlx"
)

type ItemRepository struct {
	client *sqlx.DB
}

func NewItemRepository(db *sqlx.DB) ItemRepository {
	return ItemRepository{db}
}

func (r *ItemRepository) SaveItem(item model.Item) (int, error) {
	query := `INSERT INTO items (description, amount)
	VALUES (?,?)`

	tx := r.client.MustBegin()
	tx.MustExec(query, item.Description, item.Amount)
	if err := tx.Commit(); err != nil {
		log.Println("[ItemRepository Error]", err)
		tx.Rollback()
		return 0, err
	}

	var lastId int
	err := r.client.Get(&lastId, "SELECT LAST_INSERT_ID()")
	if err != nil {
		log.Println("[ItemRepository Error]", err)
		return 0, err
	}

	return lastId, nil

}

func (r *ItemRepository) SaveItemInvoice(itemID, invoiceID int) error {
	query := `INSERT INTO item_invoice (item_id, invoice_item)
	VALUES (?,?)`

	tx := r.client.MustBegin()
	tx.MustExec(query, itemID, invoiceID)
	if err := tx.Commit(); err != nil {
		log.Println("[ItemRepository Error]", err)
		tx.Rollback()
		return err
	}

	return nil
}

func (r *ItemRepository) UpdateItem(i model.Item, id int) error {
	query := `UPDATE items SET
	description = ?,
	amount = ?
	WHERE id = ?`
	tx := r.client.MustBegin()
	tx.MustExec(query, i.Description, i.Amount, id)
	if err := tx.Commit(); err != nil {
		log.Println("[ItemRepository Error]", err)
		tx.Rollback()
		return err
	}

	return nil
}
