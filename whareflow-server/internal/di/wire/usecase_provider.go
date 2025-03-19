//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/Miroslovelife/whareflow/internal/config"
	"github.com/Miroslovelife/whareflow/internal/repositories"
	"github.com/Miroslovelife/whareflow/internal/services"
	"github.com/Miroslovelife/whareflow/internal/usecase"
	"github.com/Miroslovelife/whareflow/pkg/qr"
	"github.com/google/wire"
)

type ProviderUsecase struct {
	UserUsecase       *usecase.IUserUsecase
	WareHouseUsecase  *usecase.IWarehouseUsecase
	ZoneUsecase       *usecase.IZoneUsecase
	ProductUsecase    *usecase.IProductUsecase
	PermissionUsecase *usecase.IPermissionUsecase
	AuthUsecase       *usecase.IAuthUsecase
}

func ProvideUserUsecase(repoUser repositories.UserRepository, passwordHasher services.PasswordHasher, tokenManager services.TokenManager) *usecase.IUserUsecase {
	return usecase.NewUserUsecase(repoUser, passwordHasher, tokenManager)
}

func ProvideWarehouseUsecase(repoWarehouse repositories.WareHouseRepository) *usecase.IWarehouseUsecase {
	return usecase.NewIWarehouseUsecase(repoWarehouse)
}

func ProvideZoneUsecase(repoZone repositories.ZoneRepository) *usecase.IZoneUsecase {
	return usecase.NewIZoneUsecase(repoZone)
}

func ProvideProductUsecase(repoProduct repositories.ProductRepository, qr qr.GeneratorQR, cfg config.Config) *usecase.IProductUsecase {
	return usecase.NewIProductUsecase(repoProduct, qr, cfg)
}

func ProvidePermissionUsecase(repoUser repositories.UserRepository, repoPermission repositories.PermissionRepository, repoWarehouse repositories.WareHouseRepository) *usecase.IPermissionUsecase {
	return usecase.NewIPermissionUsecase(repoUser, repoPermission, repoWarehouse)
}

func ProvideAuthUsecase(repoUser repositories.UserRepository, tokenManager services.TokenManager) *usecase.IAuthUsecase {
	return usecase.NewIAuthUsecase(repoUser, tokenManager)
}

var UsecaseProviderSet = wire.NewSet(
	ProvideUserUsecase,
	ProvideWarehouseUsecase,
	ProvideZoneUsecase,
	ProvideProductUsecase,
	ProvidePermissionUsecase,
	ProvideAuthUsecase,
	wire.Struct(new(ProviderUsecase), "UserUsecase", "WareHouseUsecase", "ZoneUsecase", "ProductUsecase", "PermissionUsecase", "AuthUsecase"),
)

func InitializeUsecaseProviderSet(repoUser repositories.UserRepository,
	passwordHasher services.PasswordHasher,
	tokenManager services.TokenManager,
	repoWarehouse repositories.WareHouseRepository,
	repoZone repositories.ZoneRepository,
	repoProduct repositories.ProductRepository,
	qr qr.GeneratorQR,
	cfg config.Config,
	repoPermission repositories.PermissionRepository,
) ProviderUsecase {
	wire.Build(UsecaseProviderSet)
	return ProviderUsecase{}
}
