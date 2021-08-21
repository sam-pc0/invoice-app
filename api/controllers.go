package api

import (
	"github.com/canxega/invoice-app/db"
	"github.com/canxega/invoice-app/handlers"
	"github.com/canxega/invoice-app/repository"
	"github.com/canxega/invoice-app/service"
	"github.com/gorilla/mux"
)

var (
	client      = db.NewSqlClient()
	billHandler = handlers.BillHandler{S: service.NewBillService(repository.NewBillRepository(client.DB))}
)

func Controllers() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/bills", billHandler.NewBillBasic).Methods("POST")

	router.HandleFunc("/bills/{id}", billHandler.GetBillByID).Methods("GET")

	return router
}
