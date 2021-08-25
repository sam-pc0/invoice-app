package api

import (
	"github.com/canxega/invoice-app/db"
	"github.com/canxega/invoice-app/handlers"
	"github.com/canxega/invoice-app/repository"
	"github.com/canxega/invoice-app/service"
	"github.com/gorilla/mux"
)

var (
	client       = db.NewSqlClient()
	billHandler  = handlers.BillHandler{S: service.NewBillService(repository.NewBillRepository(client.DB))}
)

func Controllers() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/bills", billHandler.GetBills).Methods("GET")
	router.HandleFunc("/bills", billHandler.NewBillBasic).Methods("POST")
	router.HandleFunc("/bills/bill/{id}", billHandler.GetBillContent).Methods("GET")
	router.HandleFunc("/bills/bill/{id}", billHandler.BillUpdateContent).Methods("PUT")
	router.HandleFunc("/bills", billHandler.BillsCreate).Methods("PUT")

	return router
}
