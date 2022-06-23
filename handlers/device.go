package handlers

import (
	"net/http"
	"log"

	"github.com/sam-pc0/invoice-app/db"
	"github.com/google/uuid"
)

type Device struct {
	id string `json:"id"`
	attempts string `json:"attempts"`
	blocked_time string `json:"blocked_time"`
}

func DeviceHandler(w http.ResponseWriter, r *http.Request) {
	client := db.NewSqlClient()
	var device Device
	device.id = uuid.New().String();
	query := `INSERT INTO device (id, attempts, blocked_time)	VALUES (?, ?, ?)`
	tx := client.MustBegin()
	tx.MustExec(query, device.id, device.attempts, device.blocked_time)
	if err := tx.Commit(); err != nil {
		log.Println("[Device Error]", err)
		tx.Rollback()
		writeResponse(w, http.StatusInternalServerError, "failed")
		return
	}

	writeResponse(w, http.StatusAccepted, device.id)
}
