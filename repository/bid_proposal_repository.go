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

func (r *BidProposalRepository) SaveBidProposal(b model.BidProposal) {
	query := `INSERT INTO bid_proposal (number_bid, specifications_stimates, not_included, totalSum, withdrawn_days, withdrawn_date)
	VALUES (?, ?, ?, ?, ?, ?)`

	tx := r.client.MustBegin()
	tx.MustExec(query, b.Number, b.SpecificationStimates, b.NotIncluded, b.TotalSum, b.WithdrawnDays, b.WithdrawnDate)
	if err := tx.Commit(); err != nil {
		log.Println("[BidProposalRepository Error]", err)
		tx.Rollback()
		return
	}
}
