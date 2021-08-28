package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"

	"github.com/canxega/invoice-app/db"
	"github.com/canxega/invoice-app/model"
	"github.com/canxega/invoice-app/service"
	"github.com/canxega/invoice-app/repository"
)


var (
	client = db.NewSqlClient()
	BillService = service.NewBillService(repository.NewBillRepository(client.DB))
	InvoiceService = service.NewInvoiceService(repository.NewInvoiceRepository(client.DB)) 
	BidService = service.NewBidService(repository.NewBidProposalRepository(client.DB))
	OwnerService = service.NewOwnerService(repository.NewOwnerRepository(client.DB)) 
)

type BillHandler struct {}

func (h *BillHandler) GetBills(w http.ResponseWriter, r *http.Request) {
	bills, err := BillService.GetAllBillS()
	
	if err != nil {
		log.Println("[Handler Bill Error]", err)
		writeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeResponse(w, http.StatusOK, bills)
}

func (h *BillHandler) CreateBill(w http.ResponseWriter, r *http.Request) {
	var b model.Bill

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		log.Println("[Handler Bill Error]", err)
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	ownerId, err := OwnerService.CreateOwner(b.Owner)
	if err != nil {
		log.Println("[Handler Bill Error]", err)
		writeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	billId, err := BillService.CreateBill(ownerId, b)
	if err != nil {
		log.Println("[Handler Bill Error]", err)
		writeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	switch b.TemplateCode {
	case model.INVOICE:
		_, err := InvoiceService.CreateInvoice(billId) 
		if err != nil {
			log.Println("[Handler Bill Error]", err)
			writeResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	case model.BID_PROPOSAL:
		_, err := BidService.CreateBid(billId) 
		if err != nil {
			log.Println("[Handler Bill Error]", err)
			writeResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	default:
		writeResponse(w, http.StatusInternalServerError, "Invalid Data")
	}

	writeResponse(w, http.StatusOK, billId)
}

func (h *BillHandler) DeleteBill(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	billId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("[Handler Bill Error]", err, billId)
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	code, _ := BillService.GetTemplateCode(billId)
	switch code {
	case model.INVOICE: // Invoice
		err = InvoiceService.DeleteByBillId(billId)
		if err != nil {
			log.Println("[Handler Bill Error]", err)
			writeResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	case model.BID_PROPOSAL: // BID
	  err := BidService.DeleteByBillId(billId)
		if err != nil {
			log.Println("[Handler Bill Error]", err)
			writeResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	default:
		writeResponse(w, http.StatusInternalServerError, "Invalid Data")
	}

	err = BillService.DeleteById(billId)
	if err != nil {
		log.Println("[Handler Bill Error]", err)
		writeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = OwnerService.R.DeleteOwnerbyBillId(billId)
	if err != nil {
		log.Println("[Handler Bill Error]", err)
		writeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeResponse(w, http.StatusOK, "success")
}

func (h *BillHandler) GetBillById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("[Handler Bill Error]", err, id)
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	code, _ := BillService.GetTemplateCode(id)
	switch code {
	case model.INVOICE: // Invoice
		b, err := InvoiceService.GetFullInvoiceByBillId(id)
		if err != nil {
			log.Println("[Handler Bill Error]", err)
			writeResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeResponse(w, http.StatusOK, b)
	case model.BID_PROPOSAL: // BID
		b, err := BidService.CreateBid(id)
		if err != nil {
			log.Println("[Handler Bill Error]", err)
			writeResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeResponse(w, http.StatusOK, b)
		return
	}
}

func (h *BillHandler) UpdateBill(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _:= strconv.Atoi(vars["id"])

	var b model.BillJoinInvoice
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		log.Println("[Handler Bill Error]", err)
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	owner := b.Owner
	_, err = OwnerService.R.UpdateOwnerByBillId(id, owner)
	if err != nil {
		log.Println("[Handler Bill Error]", err)
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	bill := model.Bill{
		ID:          b.ID,
		Name:        b.Name,
		Description: b.Description,
		LastEdit:    b.LastEdit,
	}

	err = BillService.UpdateBill(id, bill)
	if err != nil {
		log.Println("[Handler Bill Error]", err)
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	code, err := BillService.GetTemplateCode(id)
	if err != nil {
		log.Println("[Handler Bill Error]", err)
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	switch code {
	case 1110:
		items := b.Items
		invoice := model.Invoice{
			ID:             b.ID,
			Total:          b.Total,
			DateSubmmitted: b.DateSubmmitted,
		}

		err = InvoiceService.UpdateInvoice(id, invoice, items)
		if err != nil {
			log.Println("[Handler Bill Error]", err)
			writeResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		writeResponse(w, http.StatusAccepted, "success")
		return
	case 1100:
		var b model.BillJoinBid
		err := json.NewDecoder(r.Body).Decode(&b)
		if err != nil {
			log.Println("[Handler Bill Error]", err)
			writeResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		bid := model.BidProposal{
			ID:                    b.ID_BID,
			SpecificationStimates: b.SpecificationStimates,
			NotIncluded:           b.NotIncluded,
			TotalSum:              b.TotalSum,
			WithdrawnDays:         b.WithdrawnDays,
			WithdrawnDate:         b.WithdrawnDate,
		}
		err = BidService.UpdateBid(id, bid)
		if err != nil {
			log.Println("[Handler Bill Error]", err)
			writeResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeResponse(w, http.StatusAccepted, "success")
		return
	default:
		writeResponse(w, http.StatusBadRequest, "Invalid data")
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
