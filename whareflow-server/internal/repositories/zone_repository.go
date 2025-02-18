package repositories

import "github.com/Miroslovelife/whareflow/internal/domain"

type ZoneRepository interface {
	InsertZoneData(zone *domain.Zone) error
}
