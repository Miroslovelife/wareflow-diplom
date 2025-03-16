//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/Miroslovelife/whareflow/internal/repositories"
	"github.com/Miroslovelife/whareflow/pkg/database"
	"github.com/google/wire"
	"log/slog"
)

type ProviderRepository struct {
	UserRepo       *repositories.UserPostgresRepository
	ProductRepo    *repositories.ProductPostgresRepository
	WareHouseRepo  *repositories.WareHousePostgresRepository
	ZoneRepo       *repositories.ZonePostgresRepository
	PermissionRepo *repositories.PermissionPostgresRepository
}

// Providers for repositories

func ProvideUserRepository(db database.Database, logger slog.Logger) *repositories.UserPostgresRepository {
	return repositories.NewUserPostgresRepository(db, logger)
}

func ProvideProductRepository(db database.Database, logger slog.Logger) *repositories.ProductPostgresRepository {
	return repositories.NewProductPostgresRepository(db, logger)
}

func ProvideWareHouseRepository(db database.Database, logger slog.Logger) *repositories.WareHousePostgresRepository {
	return repositories.NewWarehouseRepository(db, logger)
}

func ProvideZoneRepository(db database.Database, logger slog.Logger) *repositories.ZonePostgresRepository {
	return repositories.NewZoneRepository(db, logger)
}

func ProvidePermissionRepository(db database.Database, logger slog.Logger) *repositories.PermissionPostgresRepository {
	return repositories.NewPermissionPostgresRepository(db, logger)
}

// RepositoryProviderSet for repo layer
var RepositoryProviderSet = wire.NewSet(
	ProvideUserRepository,
	ProvideProductRepository,
	ProvideWareHouseRepository,
	ProvideZoneRepository,
	ProvidePermissionRepository,
	wire.Struct(new(ProviderRepository), "UserRepo", "ProductRepo", "WareHouseRepo", "ZoneRepo", "PermissionRepo"),
)

func InitializeRepoProviderSet(db database.Database, logger slog.Logger) ProviderRepository {
	wire.Build(RepositoryProviderSet)
	return ProviderRepository{}
}
