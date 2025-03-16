package repositories

import (
	"errors"
	"github.com/Miroslovelife/whareflow/internal/domain"
	custom_errors "github.com/Miroslovelife/whareflow/internal/errors"
	"github.com/Miroslovelife/whareflow/pkg/database"
	"gorm.io/gorm"
	"log/slog"
)

type WareHouseRepository interface {
	InsertWareHouseData(warehouse *domain.WareHouse) error
	UpdateWareHouseData(warehouse *domain.WareHouse, warehouseId uint) error
	DeleteWareHouseData(uuid string, id uint) error
	FindAllWareHouseData(uuid string) (*[]domain.WareHouse, error)
	FindWareHouseData(uuid string, id uint) (*domain.WareHouse, error)
	FindWareHouseOwner(warehouseId uint) (*domain.WareHouse, error)
}

type WareHousePostgresRepository struct {
	db     database.Database
	logger slog.Logger
}

func NewWarehouseRepository(db database.Database, logger slog.Logger) *WareHousePostgresRepository {
	return &WareHousePostgresRepository{
		db:     db,
		logger: logger,
	}
}

func (wr *WareHousePostgresRepository) InsertWareHouseData(warehouse *domain.WareHouse) error {
	user := domain.User{}
	resultUser := wr.db.GetDb().Where("uuid = ?", warehouse.UuidUser).First(&user)
	if resultUser.Error != nil {
		return resultUser.Error
	}

	if err := wr.db.GetDb().Where("uuid_user = ? AND name = ?", warehouse.UuidUser, warehouse.Name).First(&warehouse).Error; err == nil {
		return custom_errors.ErrWarehouseAlreadyExist
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	resultWarehouse := wr.db.GetDb().Create(warehouse)

	if resultWarehouse.Error != nil {
		return resultWarehouse.Error
	}

	return nil
}

func (wr *WareHousePostgresRepository) UpdateWareHouseData(warehouse *domain.WareHouse, warehouseId uint) error {
	wh := domain.WareHouse{}

	resultUser := wr.db.GetDb().Model(&domain.User{}).Where("uuid = ?", warehouse.UuidUser)
	if resultUser.Error != nil {
		return resultUser.Error
	}

	var countRows int64

	wr.db.GetDb().Model(&wh).Where("uuid_user = ? AND name = ?", warehouse.UuidUser, warehouse.Name).Count(&countRows)
	if countRows > 1 {
		return custom_errors.ErrWarehouseAlreadyExist
	}

	resultWarehouse := wr.db.GetDb().Model(&wh).Where("uuid_user = ? AND id = ?", warehouse.UuidUser, warehouseId).Updates(warehouse)
	if resultWarehouse.Error != nil {
		return resultWarehouse.Error
	}
	if resultWarehouse.RowsAffected == 0 {
		return custom_errors.ErrWareHouseNotFound
	}

	return nil
}

func (wr *WareHousePostgresRepository) DeleteWareHouseData(uuid string, id uint) error {
	warehouse := domain.WareHouse{}

	resultErr := wr.db.GetDb().Where("id = ? AND uuid_user = ?", id, uuid).Delete(warehouse)
	if resultErr.Error != nil {
		return resultErr.Error
	}

	return nil
}

func (wr *WareHousePostgresRepository) FindAllWareHouseData(uuid string) (*[]domain.WareHouse, error) {
	var warehouses []domain.WareHouse

	resultWarehouses := wr.db.GetDb().Where("uuid_user = ?", uuid).Find(&warehouses)
	if resultWarehouses.Error != nil {
		return nil, resultWarehouses.Error
	}

	return &warehouses, nil
}

func (wr *WareHousePostgresRepository) FindWareHouseData(uuid string, id uint) (*domain.WareHouse, error) {
	var warehouse domain.WareHouse

	resultWarehouses := wr.db.GetDb().Where("id = ? AND uuid_user = ?", id, uuid).Find(&warehouse)
	if resultWarehouses.Error != nil {
		return nil, resultWarehouses.Error
	}

	return &warehouse, nil
}

func (wr *WareHousePostgresRepository) FindWareHouseOwner(warehouseId uint) (*domain.WareHouse, error) {
	var warehouse domain.WareHouse

	err := wr.db.GetDb().Where("id = ?", warehouseId).First(&warehouse).Error
	if err != nil {
		return nil, err
	}

	return &warehouse, nil
}
