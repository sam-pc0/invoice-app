package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/canxega/invoice-app/model"
	"github.com/canxega/invoice-app/service"
	"github.com/gorilla/mux"
)

type BillHandler struct {
	S service.BillServie
}

func (h *BillHandler) NewBillBasic(w http.ResponseWriter, r *http.Request) {
	var b model.Bill

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		log.Println("[Handler Bill Error]", err)
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if b.Template.TemplateCode == 0 {
		log.Println("[Handler Bill Error]", "Invalid data", b.Template.TemplateCode)
		writeResponse(w, http.StatusBadRequest, "Invalid data")
		return
	}

	id, err := h.S.SaveBasicBill(b)
	if err != nil {
		log.Println("[Handler Bill Error]", err)
		writeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeResponse(w, http.StatusCreated, id)
}

func (h *BillHandler) GetBillByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("[Handler Bill Error]", err, id)
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	b, err := h.S.GetBillById(id)
	if err != nil {
		log.Println("[Handler Bill Error]", err)
		writeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeResponse(w, http.StatusOK, b)
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
