package handlers

import (
	"net/http"

	"github.com/sam-pc0/invoice-app/db"
	"github.com/sam-pc0/invoice-app/service"
	"github.com/sam-pc0/invoice-app/repository"
)

var (
	deviceClient = db.NewSqlClient()
	DeviceService = service.NewDeviceService(repository.NewDeviceRepository(deviceClient.DB))
)

type DeviceHandler struct {}

func  (handler *DeviceHandler) CreateDevice(w http.ResponseWriter, r *http.Request) {
	deviceId, err := DeviceService.CreateDevice()
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, "failed")
	}
	writeResponse(w, http.StatusAccepted, deviceId)
}
