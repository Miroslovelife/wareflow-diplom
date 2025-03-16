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

// RepositoryProviderSet for repo layer
var HandlerProviderSet = wire.NewSet(
	ProvideUserHandler,
	ProvideWareHouseHandler,
	ProvideZoneHandler,
	ProvideProductHandler,
	wire.Struct(new(ProviderHandler), "UserHandler", "WareHouseHandler", "ZoneHandler", "ProductHandler"),
)

func InitializeHandlerProviderSet(logger slog.Logger, userUsecase usecase.UserUsecase, whUsecase usecase.WarehouseUsecase, zoneUsecase usecase.ZoneUsecase, productUsecase usecase.ProductUsecase, cfg config.Config) ProviderHandler {
	wire.Build(HandlerProviderSet)
	return ProviderHandler{}
}
