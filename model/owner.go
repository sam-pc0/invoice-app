package model

type Owner struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	Location           string `json:"location"`
	Phone              string `json:"phone"`
	AltPhone           string `json:"altphone"`
	ProjectNameAddress string `json:"project_name_address"`
	Email              string `json:"email"`
}
