package model

type LoginData struct {
	Name     string `json:"username"`
	Password string `json:"password"`
	DeviceId string `json:"deviceId"`
}
