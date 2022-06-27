package api

import (
	"github.com/sam-pc0/invoice-app/handlers"
	"github.com/gorilla/mux"
)

var (
	loginHanlder = handlers.LoginHandler{}
	billHandler = handlers.BillHandler{}
	deviceHandler = handlers.DeviceHandler{}
)

func Controllers() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/login", loginHanlder.VerifyUser).Methods("POST") //login
	router.HandleFunc("/device/create", deviceHandler.CreateDevice).Methods("GET") //create device
	router.HandleFunc("/bills", billHandler.GetBills).Methods("GET") // get all invoices
	router.HandleFunc("/bills", billHandler.CreateBill).Methods("POST") // create invoice
	router.HandleFunc("/bills/bill/{id}/", billHandler.DeleteBill).Methods("DELETE") //delete invoice
	router.HandleFunc("/bills/bill/{id}", billHandler.GetBillById).Methods("GET") // get by id
	router.HandleFunc("/bills/bill/{id}", billHandler.UpdateBill).Methods("PUT") // update
	return router
}
