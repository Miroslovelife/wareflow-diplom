package repositories

import "github.com/Miroslovelife/whareflow/internal/domain"

type ZoneRepository interface {
	InsertZoneData(zone *domain.Zone, userId string) error
	UpdateZoneData(zone *domain.Zone, userId string) error
	FindAllZoneData(userId string, warehouseId int) (*[]domain.Zone, error)
	FindZoneData(userId string, warehouseId, zoneId int) (*domain.Zone, error)
	DeleteZoneData(userId string, warehouseId, zoneId int) error
}
