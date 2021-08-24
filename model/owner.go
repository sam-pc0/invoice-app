package model

type Owner struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	Location           string `json:"location"`
	Address            string `json:"address"`
	Phone              string `json:"phone"`
	AltPhone           string `json:"altPhone" db:"altPhone"`
	ProjectNameAddress string `json:"projectNameNAddress" db:"projectNameAddress"`
	Email              string `json:"email"`
}
