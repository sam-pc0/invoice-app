package service

import (
	"log"
	"time"

	"github.com/sam-pc0/invoice-app/model"
	"github.com/sam-pc0/invoice-app/repository"
	"github.com/google/uuid"
)

type DeviceService interface {
	CreateDevice() (string, error)
	IsDeviceBlocked(bool, string) (error)
	Clear(string) (error)
	NeedToBlock(string) (bool, error)
	Block(string) (error)
	IncreaseAttempts(string) (error)
}

type DefaultDeviceService struct {
	Repo repository.DeviceRepository
}

func NewDeviceService(repo repository.DeviceRepository) DefaultDeviceService {
	return DefaultDeviceService{Repo: repo}

}

const TIME_LAYOUT = time.UnixDate
const BLOCK_MINUTES = 3 
const MAX_ATTEMPTS = 5 

func (defaultService DefaultDeviceService) CreateDevice()  (string, error) {
	now := time.Now()
	var device model.Device 
	device.Id = uuid.New().String();
	device.Attempts = 0
	blockedTime := now.Add(-time.Hour * 3)
	device.BlockedTime = blockedTime.String()
	err := defaultService.Repo.CreateDevice(device)
	if err != nil {
		log.Println("[DeviceService Error]", err)
		return "", err
	}
	return device.Id, nil 
}

func (defaultService DefaultDeviceService) IsDeviceBlocked(deviceId string)  (bool, error) {
	device, err := defaultService.Repo.GetDevice(deviceId)
	if err != nil {
		log.Println("[DeviceService Error]", err)
		return true, err
	}

	blockedTime, _ := time.Parse(TIME_LAYOUT, device.BlockedTime)
	since := int(time.Since(blockedTime).Minutes())
	return since < BLOCK_MINUTES, nil
}

func (defaultService DefaultDeviceService) Clear(deviceId string)  (error) {
	err := defaultService.Repo.Clear(deviceId)
	if err != nil {
		log.Println("[DeviceService Error]", err)
		return err
	}
	return nil
}

func (defaultService DefaultDeviceService) NeedToBlock(deviceId string)  (bool, error) {
	device, err := defaultService.Repo.GetDevice(deviceId)
	if err != nil {
		log.Println("[DeviceService Error]", err)
		return true, err
	}

	if device.Attempts + 1 >= MAX_ATTEMPTS {
		return true, nil
	}

	return false, nil
}

func (defaultService DefaultDeviceService) IncreaseAttempts(deviceId string)  (error) {
	device, err := defaultService.Repo.GetDevice(deviceId)
	if err != nil {
		log.Println("[DeviceService Error]", err)
		return  err
	}

	err = defaultService.Repo.IncreaseAttempts(device.Attempts + 1, deviceId)
	if err != nil {
		log.Println("[DeviceService Error]", err)
		return  err
	}

	return nil
}

func (defaultService DefaultDeviceService) Block(deviceId string)  (error) {
	err := defaultService.IncreaseAttempts(deviceId)
	if err != nil {
		log.Println("[DeviceService Error]", err)
		return  err
	}

	now := time.Now()
	date := now.Format(TIME_LAYOUT) 
	err = defaultService.Repo.SetBlockTime(date, deviceId)
	if err != nil {
		log.Println("[DeviceService Error]", err)
		return  err
	}

	return nil
}
