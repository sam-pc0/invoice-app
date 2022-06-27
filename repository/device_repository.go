package repository

import (
	"log"

	"github.com/sam-pc0/invoice-app/model"
	"github.com/jmoiron/sqlx"
)

type DeviceRepository struct {
	client *sqlx.DB
}

func NewDeviceRepository(db *sqlx.DB) DeviceRepository {
	return DeviceRepository{db}
}

func (repo *DeviceRepository) GetDevice(deviceId string) (model.Device, error) {
	var device model.Device
	err := repo.client.Get(&device, "SELECT * FROM device WHERE id=?", deviceId)
	if err != nil {
		log.Println("[Device Repository Error]", err)
		return model.Device{}, err 
	}
	return device, nil
}

func (repo *DeviceRepository) CreateDevice(device model.Device) (error) {
	query := `INSERT INTO device (id, attempts, blocked_time)	VALUES (?, ?, ?)`
	tx := repo.client.MustBegin()
	tx.MustExec(query, device.Id, device.Attempts, device.BlockedTime)
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		log.Println("[Device Repository Error]", err)
		return err;
	}
	return nil;
}

func (repo *DeviceRepository) Clear(deviceId string) (error) {
	query := `UPDATE device SET attempts=?, blocked_time=? WHERE id=?`
	tx := repo.client.MustBegin()
	tx.MustExec(query, 0, "0001-01-01 00:00:00 +0000 UTC", deviceId)
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		log.Println("[Device Repository Error]", err)
		return err;
	}
	return nil;
}


func (repo *DeviceRepository) IncreaseAttempts(attempts int, deviceId string) (error) {
	query := `UPDATE device SET attempts=? WHERE id=?`
	tx := repo.client.MustBegin()
	tx.MustExec(query, attempts, deviceId)
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		log.Println("[Device Repository Error]", err)
		return err;
	}
	return nil;
}

func (repo *DeviceRepository) SetBlockTime(blockedTime string, deviceId string) (error) {
	query := `UPDATE device SET blocked_time=? WHERE id=?`
	tx := repo.client.MustBegin()
	tx.MustExec(query, blockedTime, deviceId)
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		log.Println("[Device Repository Error]", err)
		return err;
	}
	return nil;
}
