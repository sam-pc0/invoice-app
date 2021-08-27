package api

import (
	"github.com/canxega/invoice-app/handlers"
	"github.com/gorilla/mux"
)

var (
	billHandler  = handlers.BillHandler{}
)

func Controllers() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST") //login
	router.HandleFunc("/bills", billHandler.GetBills).Methods("GET") // get all invoices
	router.HandleFunc("/bills", billHandler.CreateBill).Methods("POST") // create invoice
	router.HandleFunc("/bills/bill/{id}/", billHandler.DeleteBill).Methods("DELETE") //delete invoice
	router.HandleFunc("/bills/bill/{id}", billHandler.GetBillById).Methods("GET") // get by id
	router.HandleFunc("/bills/bill/{id}", billHandler.UpdateBill).Methods("PUT") // update
	return router
}
