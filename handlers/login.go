package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sam-pc0/invoice-app/db"
)

type Login struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, request *http.Request) {
	c := db.NewSqlClient()
	var l Login
	_ = json.NewDecoder(request.Body).Decode(&l)

	var login Login
	c.Get(&login, "SELECT * FROM login WHERE name=? and password=?", l.Name, l.Password)

	writeResponse(w, http.StatusAccepted, "succes")
}
