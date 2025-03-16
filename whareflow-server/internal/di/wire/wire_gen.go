// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/Miroslovelife/whareflow/internal/config"
	"github.com/Miroslovelife/whareflow/internal/deliviry/http/v1/handler"
	"github.com/Miroslovelife/whareflow/internal/deliviry/http/v1/middleware"
	"github.com/Miroslovelife/whareflow/internal/repositories"
	"github.com/Miroslovelife/whareflow/internal/services"
	"github.com/Miroslovelife/whareflow/internal/usecase"
	"github.com/Miroslovelife/whareflow/pkg/database"
	"github.com/Miroslovelife/whareflow/pkg/qr"
	"github.com/google/wire"
	"log/slog"
)

// Injectors from handler_provider.go:

func InitializeHandlerProviderSet(logger slog.Logger, userUsecase usecase.UserUsecase, whUsecase usecase.WarehouseUsecase, zoneUsecase usecase.ZoneUsecase, productUsecase usecase.ProductUsecase, cfg config.Config) ProviderHandler {
	iUserHttpHandler := ProvideUserHandler(logger, userUsecase, cfg)
	iWareHouseHandler := ProvideWareHouseHandler(logger, whUsecase, cfg)
	iZoneHandler := ProvideZoneHandler(logger, zoneUsecase, cfg)
	iProductHandler := ProvideProductHandler(productUsecase, cfg)
	providerHandler := ProviderHandler{
		UserHandler:      iUserHttpHandler,
		WareHouseHandler: iWareHouseHandler,
		ZoneHandler:      iZoneHandler,
		ProductHandler:   iProductHandler,
	}
	return providerHandler
}

// Injectors from middleware_provider.go:

func InitializeMiddlewareProviderSet(authUsecase usecase.AuthUsecase, cfg config.Config, userUsecase usecase.UserUsecase, permissionUsecase usecase.PermissionUsecase) MiddlewareProvider {
	authHttpMiddleware := AuthMiddlewareProvider(authUsecase, cfg)
	roleHttpMiddleware := RoleMiddlewareProvider(userUsecase)
	iWhPermissionMiddleware := PermissionMiddlewareProvider(permissionUsecase)
	middlewareProvider := MiddlewareProvider{
		AuthMiddleware: authHttpMiddleware,
		RoleMiddleware: roleHttpMiddleware,
		WhMiddleware:   iWhPermissionMiddleware,
	}
	return middlewareProvider
}

// Injectors from repository_provider.go:

func InitializeRepoProviderSet(db database.Database, logger slog.Logger) ProviderRepository {
	userPostgresRepository := ProvideUserRepository(db, logger)
	productPostgresRepository := ProvideProductRepository(db, logger)
	wareHousePostgresRepository := ProvideWareHouseRepository(db, logger)
	zonePostgresRepository := ProvideZoneRepository(db, logger)
	permissionPostgresRepository := ProvidePermissionRepository(db, logger)
	providerRepository := ProviderRepository{
		UserRepo:       userPostgresRepository,
		ProductRepo:    productPostgresRepository,
		WareHouseRepo:  wareHousePostgresRepository,
		ZoneRepo:       zonePostgresRepository,
		PermissionRepo: permissionPostgresRepository,
	}
	return providerRepository
}

// Injectors from service_provider.go:

func InitializeServiceProviderSet(salt string, logger slog.Logger) ProviderService {
	tokenM := ProvideTokenManagerService()
	sha1Hasher := ProvideHasherService(salt)
	generator := ProvideQRService(logger)
	providerService := ProviderService{
		TokenManager: tokenM,
		Hasher:       sha1Hasher,
		QR:           generator,
	}
	return providerService
}

// Injectors from usecase_provider.go:

func InitializeUsecaseProviderSet(repoUser repositories.UserRepository, passwordHasher services.PasswordHasher, tokenManager services.TokenManager, repoWarehouse repositories.WareHouseRepository, repoZone repositories.ZoneRepository, repoProduct repositories.ProductRepository, qr2 qr.GeneratorQR, cfg config.Config, repoPermission repositories.PermissionRepository) ProviderUsecase {
	iUserUsecase := ProvideUserUsecase(repoUser, passwordHasher, tokenManager)
	iWarehouseUsecase := ProvideWarehouseUsecase(repoWarehouse)
	iZoneUsecase := ProvideZoneUsecase(repoZone)
	iProductUsecase := ProvideProductUsecase(repoProduct, qr2, cfg)
	iPermissionUsecase := ProvidePermissionUsecase(repoPermission, repoWarehouse)
	iAuthUsecase := ProvideAuthUsecase(tokenManager)
	providerUsecase := ProviderUsecase{
		UserUsecase:       iUserUsecase,
		WareHouseUsecase:  iWarehouseUsecase,
		ZoneUsecase:       iZoneUsecase,
		ProductUsecase:    iProductUsecase,
		PermissionUsecase: iPermissionUsecase,
		AuthUsecase:       iAuthUsecase,
	}
	return providerUsecase
}

// handler_provider.go:

type ProviderHandler struct {
	UserHandler      *handler.IUserHttpHandler
	WareHouseHandler *handler.IWareHouseHandler
	ZoneHandler      *handler.IZoneHandler
	ProductHandler   *handler.IProductHandler
}

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
	ProvideProductHandler, wire.Struct(new(ProviderHandler), "UserHandler", "WareHouseHandler", "ZoneHandler", "ProductHandler"),
)

// middleware_provider.go:

type MiddlewareProvider struct {
	AuthMiddleware *custom_middleware.AuthHttpMiddleware
	RoleMiddleware *custom_middleware.RoleHttpMiddleware
	WhMiddleware   *custom_middleware.IWhPermissionMiddleware
}

func AuthMiddlewareProvider(authUsecase usecase.AuthUsecase, cfg config.Config) *custom_middleware.AuthHttpMiddleware {
	return custom_middleware.NewAuthHttpMiddleware(authUsecase, cfg)
}

func RoleMiddlewareProvider(userUsecase usecase.UserUsecase) *custom_middleware.RoleHttpMiddleware {
	return custom_middleware.NewRoleHttpMiddleware(userUsecase)
}

func PermissionMiddlewareProvider(permissionUsecase usecase.PermissionUsecase) *custom_middleware.IWhPermissionMiddleware {
	return custom_middleware.NewWhPermissionMiddleware(permissionUsecase)
}

var MiddlewareProviderSet = wire.NewSet(
	AuthMiddlewareProvider,
	RoleMiddlewareProvider,
	PermissionMiddlewareProvider, wire.Struct(new(MiddlewareProvider), "AuthMiddleware", "RoleMiddleware", "WhMiddleware"),
)

// repository_provider.go:

type ProviderRepository struct {
	UserRepo       *repositories.UserPostgresRepository
	ProductRepo    *repositories.ProductPostgresRepository
	WareHouseRepo  *repositories.WareHousePostgresRepository
	ZoneRepo       *repositories.ZonePostgresRepository
	PermissionRepo *repositories.PermissionPostgresRepository
}

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
	ProvidePermissionRepository, wire.Struct(new(ProviderRepository), "UserRepo", "ProductRepo", "WareHouseRepo", "ZoneRepo", "PermissionRepo"),
)

// service_provider.go:

type ProviderService struct {
	TokenManager *services.TokenM
	Hasher       *services.SHA1Hasher
	QR           *qr.Generator
}

func ProvideTokenManagerService() *services.TokenM {
	return services.NewTokenM()
}

func ProvideHasherService(salt string) *services.SHA1Hasher {
	return services.NewSHA1Hasher(salt)
}

func ProvideQRService(logger slog.Logger) *qr.Generator {
	return qr.NewGenerator(logger)
}

var ServiceProviderSet = wire.NewSet(
	ProvideTokenManagerService,
	ProvideHasherService,
	ProvideQRService, wire.Struct(new(ProviderService), "TokenManager", "Hasher", "QR"),
)

// usecase_provider.go:

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

func ProvideProductUsecase(repoProduct repositories.ProductRepository, qr2 qr.GeneratorQR, cfg config.Config) *usecase.IProductUsecase {
	return usecase.NewIProductUsecase(repoProduct, qr2, cfg)
}

func ProvidePermissionUsecase(repoPermission repositories.PermissionRepository, repoWarehouse repositories.WareHouseRepository) *usecase.IPermissionUsecase {
	return usecase.NewIPermissionUsecase(repoPermission, repoWarehouse)
}

func ProvideAuthUsecase(tokenManager services.TokenManager) *usecase.IAuthUsecase {
	return usecase.NewIAuthUsecase(tokenManager)
}

var UsecaseProviderSet = wire.NewSet(
	ProvideUserUsecase,
	ProvideWarehouseUsecase,
	ProvideZoneUsecase,
	ProvideProductUsecase,
	ProvidePermissionUsecase,
	ProvideAuthUsecase, wire.Struct(new(ProviderUsecase), "UserUsecase", "WareHouseUsecase", "ZoneUsecase", "ProductUsecase", "PermissionUsecase", "AuthUsecase"),
)
