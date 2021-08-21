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

type OwnerHandler struct {
	S service.OwnerService
}

func (h *OwnerHandler) CreatOwner(w http.ResponseWriter, r *http.Request) {
	var o model.Owner

	err := json.NewDecoder(r.Body).Decode(&o)
	if err != nil {
		log.Println("[Handler Owner Error]", err)
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.S.SaveOwner(o)
	if err != nil {
		log.Panicln("[Handler Owner Error]", err)
		writeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeResponse(w, http.StatusCreated, id)
}

func (h *OwnerHandler) GetOwnerByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("[Handler Owner Error]", err)
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	o, err := h.S.GetOwenerByID(id)
	if err != nil {
		log.Println("[Handler Owner Error]", err)
		writeResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeResponse(w, http.StatusOK, o)
}
