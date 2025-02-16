package usecase

import (
	delivery "github.com/Miroslovelife/whareflow/internal/deliviry/http/v1/model"
	"github.com/Miroslovelife/whareflow/internal/domain"
	"github.com/Miroslovelife/whareflow/internal/repositories"
)

type WarehouseUsecase interface {
	CreateWarehouse(in delivery.WarehouseModelRequest, userId string) error
	GetAllWarehouse(userId string) ([]delivery.WarehouseModelResponse, error)
	GetWarehouse(userId, name string) (*delivery.WarehouseModelResponse, error)
	UpdateWarehouse(in delivery.WarehouseModelRequest, warehouseName, userId string) error
	DeleteWarehouse(warehouseName, userId string) error
}

type IWarehouseUsecase struct {
	warehouseRepo repositories.WareHouseRepository
}

func NewIWarehouseUsecase(warehouseRepo repositories.WareHouseRepository) *IWarehouseUsecase {
	return &IWarehouseUsecase{
		warehouseRepo: warehouseRepo,
	}
}

func (wu *IWarehouseUsecase) CreateWarehouse(in delivery.WarehouseModelRequest, userId string) error {
	wh := domain.WareHouse{
		UuidUser: userId,
		Address:  in.Address,
		Name:     in.Name,
	}

	err := wu.warehouseRepo.InsertWareHouseData(&wh)
	if err != nil {
		return err
	}

	return nil
}

func (wu *IWarehouseUsecase) GetAllWarehouse(userId string) ([]delivery.WarehouseModelResponse, error) {

	warehousesRepo, err := wu.warehouseRepo.FindAllWareHouseData(userId)
	if err != nil {
		return nil, err
	}

	var warehouses []delivery.WarehouseModelResponse

	for _, valueRepo := range *warehousesRepo {
		warehouse := delivery.WarehouseModelResponse{
			Address: valueRepo.Address,
			Name:    valueRepo.Name,
		}

		warehouses = append(warehouses, warehouse)
	}

	return warehouses, err
}

func (wu *IWarehouseUsecase) GetWarehouse(userId, name string) (*delivery.WarehouseModelResponse, error) {

	warehouseRepo, err := wu.warehouseRepo.FindWareHouseData(userId, name)
	if err != nil {
		return nil, err
	}

	return &delivery.WarehouseModelResponse{
		Address: warehouseRepo.Address,
		Name:    warehouseRepo.Name,
	}, nil
}

func (wu *IWarehouseUsecase) UpdateWarehouse(in delivery.WarehouseModelRequest, warehouseName, userId string) error {
	wh := domain.WareHouse{
		UuidUser: userId,
		Address:  in.Address,
		Name:     in.Name,
	}

	err := wu.warehouseRepo.UpdateWareHouseData(&wh, warehouseName)
	if err != nil {
		return err
	}

	return nil
}

func (wu *IWarehouseUsecase) DeleteWarehouse(warehouseName, userId string) error {

	if err := wu.warehouseRepo.DeleteWareHouseData(userId, warehouseName); err != nil {
		return err
	}

	return nil
}
