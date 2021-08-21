package model

type Owner struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	Location           string `json:"location"`
	Phone              string `json:"phone"`
	AltPhone           string `json:"altphone" db:"altPhone"`
	ProjectNameAddress string `json:"project_name_address" db:"projectNameAddress"`
	Email              string `json:"email"`
}
