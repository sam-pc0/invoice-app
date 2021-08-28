package repository

import (
	"log"

	"github.com/canxega/invoice-app/model"
	"github.com/jmoiron/sqlx"
)

type BidProposalRepository struct {
	client *sqlx.DB
}

func NewBidProposalRepository(db *sqlx.DB) BidProposalRepository {
	return BidProposalRepository{db}
}

func (r *BidProposalRepository) CreateBid(billId int, b model.BidProposal) (int, error){
	query := `INSERT INTO bid_proposal (specifications_stimates, not_included, totalSum, withdrawn_days, withdrawn_date, id_bill)
	VALUES (?, ?, ?, ?, ?, ?)`

	tx := r.client.MustBegin()
	tx.MustExec(query, b.SpecificationStimates, b.NotIncluded, b.TotalSum, b.WithdrawnDays, b.WithdrawnDate, billId)
	if err := tx.Commit(); err != nil {
		log.Println("[BidProposalRepository Error]", err)
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

func (r *BidProposalRepository) DeleteByBillId(id int) (error) {
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
func (r *BidProposalRepository) GetFullBidByBillId(billId int) (model.BillJoinBid, error) {
	query := `SELECT bills.id "id",
		bills.template_code "template_code", bills.name "name", bills.description, bills.lastEdit "lastEdit",
		owner.id "owner.id", owner.name "owner.name", owner.phone "owner.phone", owner.location "owner.location"
		,owner.address "owner.address", owner.altPhone "owner.altPhone", 
		owner.projectNameAddress "owner.projectNameAddress", owner.email "owner.email",
		bid_proposal.id "id_bid", bid_proposal.specifications_stimates, bid_proposal.not_included,
		bid_proposal.totalSum, bid_proposal.withdrawn_days, bid_proposal.withdrawn_date 
	FROM bills 
	Right JOIN owner  
	ON bills.owner_id = owner.id
	INNER JOIN bid_proposal 
	ON bills.id = bid_proposal.id_bill 
	WHERE bills.id = ?`

	var b model.BillJoinBid
	err := r.client.Get(&b, query, billId)
	if err != nil {
		log.Println("[BidProposalRepository]", err)
		return model.BillJoinBid{}, err
	}

	return b, nil
}

func (r *BidProposalRepository) UpdateBid(billId int, b model.BidProposal) error {

	query := `UPDATE bid_proposal SET
	specifications_stimates=?,
	not_included=?,
	totalSum=?,
	withdrawn_days=?,
	withdrawn_date=?
	WHERE id=?`

	tx := r.client.MustBegin()
	tx.MustExec(query, b.SpecificationStimates, b.NotIncluded, b.TotalSum, b.WithdrawnDays, b.WithdrawnDate, billId)
	if err := tx.Commit(); err != nil {
		log.Println("[BidProposalRepository Error]", err)
		tx.Rollback()
		return err
	}

	return nil
}
