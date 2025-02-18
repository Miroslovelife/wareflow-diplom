package repositories

import (
	"github.com/Miroslovelife/whareflow/internal/domain"
	"github.com/Miroslovelife/whareflow/pkg/database"
	"log/slog"
)

type zonePostgresRepository struct {
	db     database.Database
	logger slog.Logger
}

func NewZoneRepository(db database.Database, logger slog.Logger) *zonePostgresRepository {
	return &zonePostgresRepository{
		db:     db,
		logger: logger,
	}
}

func (wr *zonePostgresRepository) InsertZoneData(zone *domain.Zone) error {
	resultUser := wr.db.GetDb().Model(domain.WareHouse{}).Where("id = ?", zone.WarehouseId)
	if resultUser.Error != nil {
		return resultUser.Error
	}

	resultWarehouse := wr.db.GetDb().Create(zone)

	if resultWarehouse.Error != nil {
		return resultWarehouse.Error
	}

	return nil
}

//func (wr *wareHousePostgresRepository) UpdateZoneData(zone *domain.Zone, zoneName string) error {
//	wh := domain.WareHouse{}
//
//	resultUser := wr.db.GetDb().Model(&domain.User{}).Where("uuid = ?", warehouse.UuidUser)
//	if resultUser.Error != nil {
//		return resultUser.Error
//	}
//
//	var countRows int64
//
//	wr.db.GetDb().Model(&wh).Where("uuid_user = ? AND name = ?", warehouse.UuidUser, warehouse.Name).Count(&countRows)
//	if countRows > 1 {
//		return custom_errors.ErrWarehouseAlreadyExist
//	}
//
//	resultWarehouse := wr.db.GetDb().Model(&wh).Where("uuid_user = ? AND name = ?", warehouse.UuidUser, warehouseName).Updates(warehouse)
//	if resultWarehouse.Error != nil {
//		return resultWarehouse.Error
//	}
//	if resultWarehouse.RowsAffected == 0 {
//		return custom_errors.ErrWareHouseNotFound
//	}
//
//	return nil
//}
//
//func (wr *wareHousePostgresRepository) DeleteZoneHouseData(uuid, name string) error {
//	warehouse := domain.WareHouse{}
//
//	resultErr := wr.db.GetDb().Where("name = ? AND uuid_user = ?", name, uuid).Delete(warehouse)
//	if resultErr.Error != nil {
//		return resultErr.Error
//	}
//
//	return nil
//}
//
//func (wr *wareHousePostgresRepository) FindAllZoneHouseData(uuid string) (*[]domain.Zone, error) {
//	var warehouses []domain.WareHouse
//
//	resultWarehouses := wr.db.GetDb().Where("uuid_user = ?", uuid).Find(&warehouses)
//	if resultWarehouses.Error != nil {
//		return nil, resultWarehouses.Error
//	}
//
//	return &warehouses, nil
//}
//
//func (wr *wareHousePostgresRepository) FindZoneData(uuid, name string) (*domain.Zone, error) {
//	var warehouse domain.WareHouse
//
//	resultWarehouses := wr.db.GetDb().Where("name = ? AND uuid_user = ?", name, uuid).Find(&warehouse)
//	if resultWarehouses.Error != nil {
//		return nil, resultWarehouses.Error
//	}
//
//	return &warehouse, nil
//}
