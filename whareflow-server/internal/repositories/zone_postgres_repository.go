package repositories

import (
	"github.com/Miroslovelife/whareflow/internal/domain"
	custom_errors "github.com/Miroslovelife/whareflow/internal/errors"
	"github.com/Miroslovelife/whareflow/pkg/database"
	"log/slog"
)

type ZoneRepository interface {
	InsertZoneData(zone *domain.Zone, userId string) error
	UpdateZoneData(zone *domain.Zone, userId string) error
	FindAllZoneData(userId string, warehouseId int) (*[]domain.Zone, error)
	FindZoneData(userId string, warehouseId, zoneId int) (*domain.Zone, error)
	DeleteZoneData(userId string, warehouseId, zoneId int) error
}

type ZonePostgresRepository struct {
	db     database.Database
	logger slog.Logger
}

func NewZoneRepository(db database.Database, logger slog.Logger) *ZonePostgresRepository {
	return &ZonePostgresRepository{
		db:     db,
		logger: logger,
	}
}

func (wr *ZonePostgresRepository) InsertZoneData(zone *domain.Zone, userId string) error {
	var count int64
	wr.db.GetDb().Model(domain.WareHouse{}).Where("id = ? AND uuid_user = ?", zone.WarehouseId, userId).Count(&count)
	if count != 1 {
		return custom_errors.ErrWareHouseNotFound
	}

	resultZone := wr.db.GetDb().Create(zone)

	if resultZone.Error != nil {
		return resultZone.Error
	}

	return nil
}

func (wr *ZonePostgresRepository) UpdateZoneData(zone *domain.Zone, userId string) error {

	var count int64
	wr.db.GetDb().Model(domain.WareHouse{}).Where("id = ? AND uuid_user = ?", zone.WarehouseId, userId).Count(&count)
	if count != 1 {
		return custom_errors.ErrWareHouseNotFound
	}

	resultZone := wr.db.GetDb().Model(zone).Where("id = ?", zone.Id).Updates(zone)
	if resultZone.Error != nil {
		return resultZone.Error
	}

	return nil
}

func (wr *ZonePostgresRepository) FindAllZoneData(userId string, warehouseId int) (*[]domain.Zone, error) {
	var zones []domain.Zone
	var countWarehouse int64

	wr.db.GetDb().Model(domain.WareHouse{}).Where("id = ? AND uuid_user = ?", warehouseId, userId).Count(&countWarehouse)
	if countWarehouse != 1 {
		return nil, custom_errors.ErrWareHouseNotFound
	}

	err := wr.db.GetDb().Where("ware_house_id = ?", warehouseId).Find(&zones)
	if err.Error != nil {
		return nil, err.Error
	}

	return &zones, nil
}

func (wr *ZonePostgresRepository) FindZoneData(userId string, warehouseId, zoneId int) (*domain.Zone, error) {
	var zones domain.Zone
	var countWarehouse int64

	wr.db.GetDb().Model(domain.WareHouse{}).Where("id = ? AND uuid_user = ?", warehouseId, userId).Count(&countWarehouse)
	if countWarehouse != 1 {
		return nil, custom_errors.ErrWareHouseNotFound
	}

	err := wr.db.GetDb().Where("ware_house_id = ? AND id = ?", warehouseId, zoneId).Find(&zones)
	if err.Error != nil {
		return nil, err.Error
	}

	return &zones, nil
}

func (wr *ZonePostgresRepository) DeleteZoneData(userId string, warehouseId, zoneId int) error {
	var zone domain.Zone
	var countWarehouse int64

	wr.db.GetDb().Model(domain.WareHouse{}).Where("id = ? AND uuid_user = ?", warehouseId, userId).Count(&countWarehouse)
	if countWarehouse != 1 {
		return custom_errors.ErrWareHouseNotFound
	}

	err := wr.db.GetDb().Where("ware_house_id = ? AND id = ?", warehouseId, zoneId).Delete(&zone)
	if err.Error != nil {
		return err.Error
	}

	return nil
}
