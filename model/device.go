package model

type Device struct {
	Id string `json:"id" db:"id"`
	Attempts int `json:"attempts" db:"attempts"`
	BlockedTime string `json:"blockedTime" db:"blocked_time"`
}
