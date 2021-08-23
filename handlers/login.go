package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/canxega/invoice-app/db"
)

type Login struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	c := db.NewSqlClient()
	var l Login
	_ = json.NewDecoder(r.Body).Decode(&l)

	var login Login
	c.Get(&login, "SELEC * FROM login WHERE name=? and password=?", l.Name, l.Password)

	writeResponse(w, http.StatusAccepted, "succes")
}
