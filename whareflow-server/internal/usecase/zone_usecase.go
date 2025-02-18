package usecase

import (
	delivery "github.com/Miroslovelife/whareflow/internal/deliviry/http/v1/model"
	"github.com/Miroslovelife/whareflow/internal/domain"
	"github.com/Miroslovelife/whareflow/internal/repositories"
)

type ZoneUsecase interface {
	CreateZone(in delivery.ZoneModelRequest) error
}

type IZoneUsecase struct {
	zoneRepository repositories.ZoneRepository
}

func NewZoneUsecase(zoneRepo repositories.ZoneRepository) *IZoneUsecase {
	return &IZoneUsecase{
		zoneRepository: zoneRepo,
	}
}

func (zu *IZoneUsecase) CreateZone(in delivery.ZoneModelRequest) error {
	zone := &domain.Zone{
		Name:        in.Name,
		Capacity:    in.Capacity,
		WarehouseId: in.WarehouseId,
	}

	if err := zu.zoneRepository.InsertZoneData(zone); err != nil {
		return err
	}

	return nil
}
