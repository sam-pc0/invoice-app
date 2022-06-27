package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sam-pc0/invoice-app/db"
	"github.com/sam-pc0/invoice-app/model"
	"github.com/sam-pc0/invoice-app/repository"
	"github.com/sam-pc0/invoice-app/service"
)

var (
	loginClient = db.NewSqlClient()
	LoginService = service.NewLoginService(repository.NewLoginRepository(loginClient.DB))
)

type LoginHandler struct {}

func  (handler *LoginHandler) VerifyUser(w http.ResponseWriter, request *http.Request) {
	var loginData model.LoginData 
	_ = json.NewDecoder(request.Body).Decode(&loginData)

	isDeviceBlocked, err := DeviceService.IsDeviceBlocked(loginData.DeviceId)
	if err != nil {
		log.Println("Blocked Verification Error")
		writeResponse(w, http.StatusInternalServerError, "error")
		return
	}

	if isDeviceBlocked {
		writeResponse(w, http.StatusUnauthorized, "blocked")
		return
	}

	err = LoginService.VerifyUser(loginData)
	if err == nil {
		DeviceService.Clear(loginData.DeviceId)
		writeResponse(w, http.StatusAccepted, "succes")
		return
	}

	needToBlock, err := DeviceService.NeedToBlock(loginData.DeviceId)
	if err != nil {
		log.Println("NeedToBlock Error")
		writeResponse(w, http.StatusInternalServerError, "error")
		return
	}
	
	if needToBlock {
		err =	DeviceService.Block(loginData.DeviceId)
		if err != nil {
			log.Println("Block Error")
			writeResponse(w, http.StatusInternalServerError, "error")
			return
		}
		writeResponse(w, http.StatusUnauthorized, "blocked")
		return
	}

	err = DeviceService.IncreaseAttempts(loginData.DeviceId)
	if err != nil {
		log.Println("IncreaseAttempts Error")
		writeResponse(w, http.StatusInternalServerError, "error")
		return
	}

	writeResponse(w, http.StatusUnauthorized, "error")
}
