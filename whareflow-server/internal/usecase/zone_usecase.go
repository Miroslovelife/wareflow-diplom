package usecase

import (
	delivery "github.com/Miroslovelife/whareflow/internal/deliviry/http/v1/model"
	"github.com/Miroslovelife/whareflow/internal/domain"
	"github.com/Miroslovelife/whareflow/internal/repositories"
)

type ZoneUsecase interface {
	CreateZone(in delivery.ZoneModelRequest, userId string, warehouseId int) error
	UpdateZone(in delivery.ZoneModelRequest, userId string, zoneId int, warehouseId int) error
	GetAllZone(userId string, warehouseId int) ([]delivery.ZoneModelResponse, error)
	GetZone(userId string, warehouseId, zoneId int) (*delivery.ZoneModelResponse, error)
	DeleteZone(userId string, warehouseId, zoneId int) error
}

type IZoneUsecase struct {
	zoneRepository repositories.ZoneRepository
}

func NewZoneUsecase(zoneRepo repositories.ZoneRepository) *IZoneUsecase {
	return &IZoneUsecase{
		zoneRepository: zoneRepo,
	}
}

func (zu *IZoneUsecase) CreateZone(in delivery.ZoneModelRequest, userId string, warehouseId int) error {
	zone := &domain.Zone{
		Name:        in.Name,
		Capacity:    in.Capacity,
		WarehouseId: warehouseId,
	}

	if err := zu.zoneRepository.InsertZoneData(zone, userId); err != nil {
		return err
	}

	return nil
}

func (zu *IZoneUsecase) UpdateZone(in delivery.ZoneModelRequest, userId string, zoneId int, warehouseId int) error {
	zone := &domain.Zone{
		Id:          zoneId,
		Name:        in.Name,
		Capacity:    in.Capacity,
		WarehouseId: warehouseId,
	}

	if err := zu.zoneRepository.UpdateZoneData(zone, userId); err != nil {
		return err
	}

	return nil
}

func (zu *IZoneUsecase) GetAllZone(userId string, warehouseId int) ([]delivery.ZoneModelResponse, error) {
	zonesRepo, err := zu.zoneRepository.FindAllZoneData(userId, warehouseId)
	if err != nil {
		return nil, err
	}

	var zones []delivery.ZoneModelResponse

	for _, zonesRepoValue := range *zonesRepo {
		zone := delivery.ZoneModelResponse{
			Id:       zonesRepoValue.Id,
			Name:     zonesRepoValue.Name,
			Capacity: zonesRepoValue.Capacity,
		}

		zones = append(zones, zone)
	}

	return zones, nil
}

func (zu *IZoneUsecase) GetZone(userId string, warehouseId, zoneId int) (*delivery.ZoneModelResponse, error) {
	zonesRepo, err := zu.zoneRepository.FindZoneData(userId, warehouseId, zoneId)
	if err != nil {
		return nil, err
	}

	zone := delivery.ZoneModelResponse{
		Id:       zonesRepo.Id,
		Name:     zonesRepo.Name,
		Capacity: zonesRepo.Capacity,
	}

	return &zone, nil
}

func (zu *IZoneUsecase) DeleteZone(userId string, warehouseId, zoneId int) error {
	err := zu.zoneRepository.DeleteZoneData(userId, warehouseId, zoneId)
	if err != nil {
		return err
	}

	return nil
}
