package repositories

import (
	"errors"
	"github.com/Miroslovelife/whareflow/internal/domain"
	custom_errors "github.com/Miroslovelife/whareflow/internal/errors"
	"github.com/Miroslovelife/whareflow/pkg/database"
	"gorm.io/gorm"
	"log/slog"
)

type wareHousePostgresRepository struct {
	db     database.Database
	logger slog.Logger
}

func NewWarehouseRepository(db database.Database, logger slog.Logger) *wareHousePostgresRepository {
	return &wareHousePostgresRepository{
		db:     db,
		logger: logger,
	}
}

func (wr *wareHousePostgresRepository) InsertWareHouseData(warehouse *domain.WareHouse) error {
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

func (wr *wareHousePostgresRepository) UpdateWareHouseData(warehouse *domain.WareHouse, warehouseName string) error {
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

	resultWarehouse := wr.db.GetDb().Model(&wh).Where("uuid_user = ? AND name = ?", warehouse.UuidUser, warehouseName).Updates(warehouse)
	if resultWarehouse.Error != nil {
		return resultWarehouse.Error
	}
	if resultWarehouse.RowsAffected == 0 {
		return custom_errors.ErrWareHouseNotFound
	}

	return nil
}

func (wr *wareHousePostgresRepository) DeleteWareHouseData(uuid, name string) error {
	warehouse := domain.WareHouse{}

	resultErr := wr.db.GetDb().Where("name = ? AND uuid_user = ?", name, uuid).Delete(warehouse)
	if resultErr.Error != nil {
		return resultErr.Error
	}

	return nil
}

func (wr *wareHousePostgresRepository) FindAllWareHouseData(uuid string) (*[]domain.WareHouse, error) {
	var warehouses []domain.WareHouse

	resultWarehouses := wr.db.GetDb().Where("uuid_user = ?", uuid).Find(&warehouses)
	if resultWarehouses.Error != nil {
		return nil, resultWarehouses.Error
	}

	return &warehouses, nil
}

func (wr *wareHousePostgresRepository) FindWareHouseData(uuid, name string) (*domain.WareHouse, error) {
	var warehouse domain.WareHouse

	resultWarehouses := wr.db.GetDb().Where("name = ? AND uuid_user = ?", name, uuid).Find(&warehouse)
	if resultWarehouses.Error != nil {
		return nil, resultWarehouses.Error
	}

	return &warehouse, nil
}
