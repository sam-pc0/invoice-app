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
	ownerHandler = handlers.OwnerHandler{S: service.NewOwnerService(repository.NewOwnerRepository(client.DB))}
)

func Controllers() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/bills", billHandler.NewBillBasic).Methods("POST")
	router.HandleFunc("/owners", ownerHandler.CreatOwner).Methods("POST")

	router.HandleFunc("/bills", billHandler.GetBills).Methods("GET")
	router.HandleFunc("/bills/{id}", billHandler.GetBillByID).Methods("GET")
	router.HandleFunc("/owners/{id}", ownerHandler.GetOwnerByID).Methods("GET")

	router.HandleFunc("/bills", billHandler.BillsCreate).Methods("PUT")

	return router
}
