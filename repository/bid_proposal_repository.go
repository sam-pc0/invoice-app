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

func (r *BidProposalRepository) SaveBidProposal(b model.BidProposal, id int) {
	query := `INSERT INTO bid_proposal (number_bid, specifications_stimates, not_included, totalSum, withdrawn_days, withdrawn_date, id_bill)
	VALUES (?, ?, ?, ?, ?, ?, ?)`

	tx := r.client.MustBegin()
	tx.MustExec(query, b.Number, b.SpecificationStimates, b.NotIncluded, b.TotalSum, b.WithdrawnDays, b.WithdrawnDate, id)
	if err := tx.Commit(); err != nil {
		log.Println("[BidProposalRepository Error]", err)
		tx.Rollback()
		return
	}
}

func (r *BidProposalRepository) GetBidAndBillByID(id int) (model.BillJionBid, error) {
	query := `SELECT bills.id "id",
		bills.template_code "template_code", bills.name "name", bills.description, bills.lastEdit "lastEdit",
		owner.id "owner.id", owner.name "owner.name", owner.phone "owner.phone", owner.altPhone "owner.altPhone", 
		owner.projectNameAddress "owner.projectNameAddress", owner.email "owner.email",
		bid_proposal.id "id_bid", bid_proposal.number_bid "number_bid", bid_proposal.specifications_stimates, bid_proposal.not_included,
		bid_proposal.totalSum, bid_proposal.withdrawn_days, bid_proposal.withdrawn_date 
	FROM bills 
	Right JOIN owner  
	ON bills.owner_id = owner.id
	INNER JOIN bid_proposal 
	ON bills.id = bid_proposal.id_bill 
	WHERE bills.id = ?`

	var b model.BillJionBid
	err := r.client.Get(&b, query, id)
	if err != nil {
		log.Println("[BidProposalRepository]", err)
		return model.BillJionBid{}, err
	}

	return b, nil
}
