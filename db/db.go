package db

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySqlClient struct {
	*sqlx.DB
}

func NewSqlClient() *MySqlClient {
	connStr := "clopez:C@rlos2990@(143.208.180.250:21748)/invoiceappdb"
	db, err := sqlx.Open("mysql", connStr)
	if err != nil {
		log.Fatalln("[DB Connection Error]", err)
	}

	if err = db.Ping(); err != nil {
		log.Println("[DB Connection Error]", err)
	}

	return &MySqlClient{db}
}
