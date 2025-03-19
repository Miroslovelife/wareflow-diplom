//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/Miroslovelife/whareflow/internal/config"
	"github.com/Miroslovelife/whareflow/internal/deliviry/http/v1/handler"
	"github.com/Miroslovelife/whareflow/internal/usecase"
	"github.com/google/wire"
	"log/slog"
)

type ProviderHandler struct {
	UserHandler      *handler.IUserHttpHandler
	WareHouseHandler *handler.IWareHouseHandler
	ZoneHandler      *handler.IZoneHandler
	ProductHandler   *handler.IProductHandler
	RoleHandler      *handler.IRoleHandler
}

// Providers for repositories

func ProvideUserHandler(logger slog.Logger, userUsecase usecase.UserUsecase, cfg config.Config) *handler.IUserHttpHandler {
	return handler.NewIUserHttpHandler(logger, userUsecase, cfg)
}

func ProvideWareHouseHandler(logger slog.Logger, whUsecase usecase.WarehouseUsecase, cfg config.Config) *handler.IWareHouseHandler {
	return handler.NewIWareHouseHandler(logger, whUsecase)
}

func ProvideZoneHandler(logger slog.Logger, zoneUsecase usecase.ZoneUsecase, cfg config.Config) *handler.IZoneHandler {
	return handler.NewIZoneHandler(logger, zoneUsecase)
}

func ProvideProductHandler(productUsecase usecase.ProductUsecase, cfg config.Config) *handler.IProductHandler {
	return handler.NewIProductHandler(productUsecase)
}

func ProvideRoleHandler(permUsecase usecase.PermissionUsecase) *handler.IRoleHandler {
	return handler.NewIRolehandler(permUsecase)
}

// RepositoryProviderSet for repo layer
var HandlerProviderSet = wire.NewSet(
	ProvideUserHandler,
	ProvideWareHouseHandler,
	ProvideZoneHandler,
	ProvideProductHandler,
	ProvideRoleHandler,
	wire.Struct(new(ProviderHandler), "UserHandler", "WareHouseHandler", "ZoneHandler", "ProductHandler", "RoleHandler"),
)

func InitializeHandlerProviderSet(logger slog.Logger, userUsecase usecase.UserUsecase, whUsecase usecase.WarehouseUsecase, zoneUsecase usecase.ZoneUsecase, productUsecase usecase.ProductUsecase, cfg config.Config, permUsecase usecase.PermissionUsecase) ProviderHandler {
	wire.Build(HandlerProviderSet)
	return ProviderHandler{}
}
