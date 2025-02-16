package repositories

import "github.com/Miroslovelife/whareflow/internal/domain"

type WareHouseRepository interface {
	InsertWareHouseData(warehouse *domain.WareHouse) error
	UpdateWareHouseData(warehouse *domain.WareHouse, warehouseName string) error
	DeleteWareHouseData(uuid, name string) error
	FindAllWareHouseData(uuid string) (*[]domain.WareHouse, error)
	FindWareHouseData(uuid, name string) (*domain.WareHouse, error)
}
