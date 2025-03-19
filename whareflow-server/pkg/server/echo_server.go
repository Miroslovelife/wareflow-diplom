package server

import (
	"fmt"
	_ "github.com/Miroslovelife/whareflow/docs"
	"github.com/Miroslovelife/whareflow/internal/config"
	"github.com/Miroslovelife/whareflow/internal/deliviry/http/v1/handler"
	custom_middleware "github.com/Miroslovelife/whareflow/internal/deliviry/http/v1/middleware"
	"github.com/Miroslovelife/whareflow/internal/di/wire"
	"github.com/Miroslovelife/whareflow/pkg/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	"log/slog"
)

// @title WareFlow api
// @version 1.0

// @host localhost:8089
// @basePath /api/v1/employer

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

type echoServer struct {
	app    *echo.Echo
	db     database.Database
	logger slog.Logger
	cfg    config.Config
}

type DeliveryLayer struct {
	userHandlers         *handler.IUserHttpHandler
	warehouseHandlers    *handler.IWareHouseHandler
	zoneHandlers         *handler.IZoneHandler
	productHandlers      *handler.IProductHandler
	roleHandler          *handler.IRoleHandler
	authMiddleware       *custom_middleware.AuthHttpMiddleware
	roleMiddleware       *custom_middleware.RoleHttpMiddleware
	permissionMiddleware *custom_middleware.IWhPermissionMiddleware
}

func NewEchoServer(logger slog.Logger, db database.Database, cfg *config.Config) *echoServer {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)

	return &echoServer{
		app:    echoApp,
		db:     db,
		logger: logger,
		cfg:    *cfg,
	}
}

func (s *echoServer) Start() {
	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())

	s.app.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins:     []string{"http://127.0.0.1", "http://localhost"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Authorization", "Content-Type"},
			AllowCredentials: true,
		},
	))

	s.app.Static("/qr_storage", "./qr_storage")

	v1 := s.app.Group("api/v1")

	s.app.GET("/swagger/*", echoSwagger.WrapHandler)

	delivery := s.InitLayers()

	admin := v1.Group("/admin", delivery.roleMiddleware.IsAdmin)
	owner := v1.Group("/owner", delivery.authMiddleware.Auth, delivery.roleMiddleware.IsOwner)
	employer := v1.Group("/employer", delivery.authMiddleware.Auth, delivery.roleMiddleware.IsEmployer)

	s.InitGeneralRoutes(v1, delivery)
	s.InitAdminRoutes(admin, delivery)
	s.InitOwnerRoutes(owner, delivery)
	s.InitEmployerRoutes(employer, delivery)

	s.app.Logger.Fatal(s.app.Start(fmt.Sprintf("0.0.0.0:%d", 8089)))
}

func (s *echoServer) InitLayers() *DeliveryLayer {
	repoLayer := wire.InitializeRepoProviderSet(s.db, s.logger)

	serviceLayer := wire.InitializeServiceProviderSet(s.cfg.Auth.PasswordSalt, s.logger)

	usecaseLayer := wire.InitializeUsecaseProviderSet(
		repoLayer.UserRepo,
		serviceLayer.Hasher,
		serviceLayer.TokenManager,
		repoLayer.WareHouseRepo,
		repoLayer.ZoneRepo,
		repoLayer.ProductRepo,
		serviceLayer.QR,
		s.cfg,
		repoLayer.PermissionRepo,
	)

	handlerLayer := wire.InitializeHandlerProviderSet(
		s.logger,
		usecaseLayer.UserUsecase,
		usecaseLayer.WareHouseUsecase,
		usecaseLayer.ZoneUsecase,
		usecaseLayer.ProductUsecase,
		s.cfg,
		usecaseLayer.PermissionUsecase,
	)

	middlewareLayer := wire.InitializeMiddlewareProviderSet(
		usecaseLayer.AuthUsecase,
		s.cfg,
		usecaseLayer.UserUsecase,
		usecaseLayer.PermissionUsecase,
	)

	return &DeliveryLayer{
		userHandlers:         handlerLayer.UserHandler,
		warehouseHandlers:    handlerLayer.WareHouseHandler,
		zoneHandlers:         handlerLayer.ZoneHandler,
		productHandlers:      handlerLayer.ProductHandler,
		roleHandler:          handlerLayer.RoleHandler,
		authMiddleware:       middlewareLayer.AuthMiddleware,
		roleMiddleware:       middlewareLayer.RoleMiddleware,
		permissionMiddleware: middlewareLayer.WhMiddleware,
	}

}

func (s *echoServer) InitGeneralRoutes(group *echo.Group, delivery *DeliveryLayer) {
	userRouters := group.Group("/auth")
	userRouters.POST("/sign-up", delivery.userHandlers.Register)
	userRouters.POST("/sign-in-phone", delivery.userHandlers.LoginByPhoneNumber)
	userRouters.POST("/sign-in-email", delivery.userHandlers.LoginByEmail)
	userRouters.GET("/refresh", delivery.userHandlers.Refresh)
}

func (s *echoServer) InitAdminRoutes(group *echo.Group, delivery *DeliveryLayer) {
	warehouseRouters := group.Group("/warehouse")
	warehouseRouters.GET("", delivery.warehouseHandlers.GetAllWarehouses)
	warehouseRouters.GET("/:warehouse_id", delivery.warehouseHandlers.GetWarehouse)
	warehouseRouters.POST("", delivery.warehouseHandlers.CreateWarehouse)
	warehouseRouters.PUT("/:warehouse_id", delivery.warehouseHandlers.UpdateWarehouse)
	warehouseRouters.DELETE("/:warehouse_id", delivery.warehouseHandlers.DeleteWarehouse)

	zoneRouters := warehouseRouters.Group("/:warehouse_id/zone")
	zoneRouters.GET("", delivery.zoneHandlers.GetAllZones)
	zoneRouters.GET("/:zone_id", delivery.zoneHandlers.GetZone)
	zoneRouters.POST("", delivery.zoneHandlers.CreateZone)
	zoneRouters.PUT("/:zone_id", delivery.zoneHandlers.UpdateZone)
	zoneRouters.DELETE("/:zone_id", delivery.zoneHandlers.DeleteZone)

	productRouters := zoneRouters.Group("/:zone_id/product")
	productRouters.GET("", delivery.productHandlers.GetProduct)
	productRouters.POST("", delivery.productHandlers.CreateProduct)
}

func (s *echoServer) InitOwnerRoutes(group *echo.Group, delivery *DeliveryLayer) {
	group.GET("/product/:product_id", delivery.productHandlers.GetProduct)
	group.GET("/permission", delivery.roleHandler.GetAllPermissionTypes) // Return all permission types
	warehouseRouters := group.Group("/warehouse")
	warehouseRouters.GET("", delivery.warehouseHandlers.GetAllWarehouses)
	warehouseRouters.GET("/:warehouse_id", delivery.warehouseHandlers.GetWarehouse)
	warehouseRouters.POST("", delivery.warehouseHandlers.CreateWarehouse)
	warehouseRouters.PUT("/:warehouse_id", delivery.warehouseHandlers.UpdateWarehouse)
	warehouseRouters.DELETE("/:warehouse_id", delivery.warehouseHandlers.DeleteWarehouse)

	zoneRouters := warehouseRouters.Group("/:warehouse_id/zone")
	zoneRouters.GET("", delivery.zoneHandlers.GetAllZones)
	zoneRouters.GET("/:zone_id", delivery.zoneHandlers.GetZone)
	zoneRouters.POST("", delivery.zoneHandlers.CreateZone)
	zoneRouters.PUT("/:zone_id", delivery.zoneHandlers.UpdateZone)
	zoneRouters.DELETE("/:zone_id", delivery.zoneHandlers.DeleteZone)

	productZoneRouters := zoneRouters.Group("/:zone_id/product")
	productZoneRouters.GET("/:product_id", delivery.productHandlers.GetProduct)
	productZoneRouters.GET("", delivery.productHandlers.GetAllProductsFromZone)

	productWarehouseRouters := warehouseRouters.Group("/:warehouse_id/product")
	productWarehouseRouters.GET("", delivery.productHandlers.GetAllProductsFromWarehouse)
	productWarehouseRouters.PUT("/:product_id", delivery.productHandlers.UpdateProduct)
	productWarehouseRouters.POST("", delivery.productHandlers.CreateProduct)

	employerWarehouseRoutes := warehouseRouters.Group("")
	employerWarehouseRoutes.GET("/:warehouse_id/employer", delivery.warehouseHandlers.GetEmployers)

	//Role Management
	roleRoutes := group.Group("/role")

	roleRoutes.POST("/:warehouse_id", delivery.roleHandler.GiveRoleForEmployer)
	roleRoutes.POST("/permission/:warehouse_id", delivery.roleHandler.GetAllUserPermissionOnWh)
}

func (s *echoServer) InitEmployerRoutes(group *echo.Group, delivery *DeliveryLayer) {
	warehouseRouters := group.Group("/warehouse")
	group.GET("/warehouse", delivery.warehouseHandlers.GetWhsEmployer) // Получение списка складов для работника
	group.POST("/permission/:warehouse_id/:action", delivery.roleHandler.GetAllUserPermissionOnWh,
		delivery.permissionMiddleware.SetGroup("self_perm"),
		delivery.permissionMiddleware.HasPermissionOnWarehouse)
	group.GET("/permission", delivery.roleHandler.GetAllPermissionTypes) // Получение типов прав для работника

	warehouseRoutersManage := warehouseRouters.Group("/global/:warehouse_id/:action", delivery.permissionMiddleware.SetGroup("warehouse"),
		delivery.permissionMiddleware.HasPermissionOnWarehouse)
	warehouseRoutersManage.GET("", delivery.warehouseHandlers.GetWarehouse) // Получение информации о складе

	// Роли (работники)
	roleRoutes := warehouseRouters.Group("/role/:warehouse_id/:action",
		delivery.permissionMiddleware.SetGroup("role"),
		delivery.permissionMiddleware.HasPermissionOnWarehouse)
	roleRoutes.POST("", delivery.roleHandler.GiveRoleForEmployer)                 // Назначение роли работнику
	roleRoutes.POST("/permission", delivery.roleHandler.GetAllUserPermissionOnWh) // Получение прав работника на склад
	roleRoutes.GET("/employer", delivery.warehouseHandlers.GetEmployers)
	// Зоны на складе
	zoneRouters := warehouseRouters.Group("/:warehouse_id/zone/:action",
		delivery.permissionMiddleware.SetGroup("zone"),
		delivery.permissionMiddleware.HasPermissionOnWarehouse)
	zoneRouters.GET("", delivery.zoneHandlers.GetAllZones)            // Получение всех зон на складе
	zoneRouters.GET("/:zone_id", delivery.zoneHandlers.GetZone)       // Получение информации о зоне
	zoneRouters.POST("", delivery.zoneHandlers.CreateZone)            // Создание новой зоны
	zoneRouters.PUT("/:zone_id", delivery.zoneHandlers.UpdateZone)    // Обновление зоны
	zoneRouters.DELETE("/:zone_id", delivery.zoneHandlers.DeleteZone) // Удаление зоны

	// Продукты в зоне
	productZoneRouters := warehouseRouters.Group("/:warehouse_id/zone/:zone_id/product/:action",
		delivery.permissionMiddleware.SetGroup("product"),
		delivery.permissionMiddleware.HasPermissionOnWarehouse)
	productZoneRouters.GET("/:product_id", delivery.productHandlers.GetProduct) // Получение информации о продукте

	// Продукты на складе
	productWarehouseRouters := warehouseRouters.Group("/:warehouse_id/product/:action",
		delivery.permissionMiddleware.SetGroup("product"),
		delivery.permissionMiddleware.HasPermissionOnWarehouse)
	productWarehouseRouters.GET("", delivery.productHandlers.GetAllProductsFromWarehouse) // Получение всех продуктов на складе
	productWarehouseRouters.PUT("/:product_id", delivery.productHandlers.UpdateProduct)   // Обновление продукта на складе
	productWarehouseRouters.GET("", delivery.productHandlers.GetAllProductsFromZone)      // Создание нового продукта

}
