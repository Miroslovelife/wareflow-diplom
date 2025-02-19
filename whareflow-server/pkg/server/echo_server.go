package server

import (
	"fmt"
	_ "github.com/Miroslovelife/whareflow/docs"
	"github.com/Miroslovelife/whareflow/internal/config"
	"github.com/Miroslovelife/whareflow/internal/deliviry/http/v1/handler"
	custom_middleware "github.com/Miroslovelife/whareflow/internal/deliviry/http/v1/middleware"
	"github.com/Miroslovelife/whareflow/internal/repositories"
	"github.com/Miroslovelife/whareflow/internal/services"
	"github.com/Miroslovelife/whareflow/internal/usecase"
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
// @basePath /v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

type echoServer struct {
	app    *echo.Echo
	db     database.Database
	logger slog.Logger
	cfg    config.Config
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

	v1 := s.app.Group("v1")

	s.app.GET("/swagger/*", echoSwagger.WrapHandler)

	s.InitRoutes(v1)

	s.app.Logger.Fatal(s.app.Start(fmt.Sprintf(":%d", 8089)))
}

func (s *echoServer) InitRoutes(group *echo.Group) {
	// Layers
	userPostgresRepository := repositories.NewUserPostgresRepository(s.db, s.logger)
	passwordHasher := services.NewSHA1Hasher(s.cfg.Auth.PasswordSalt)
	manager := services.NewTokenM()
	userUsecase := usecase.NewUserUsecase(userPostgresRepository, passwordHasher, manager)
	userHttpHandler := handler.NewUserHttpHandler(s.logger, userUsecase, s.cfg)
	roleMiddleware := custom_middleware.NewRoleHttpMiddleware(userUsecase)

	whRepository := repositories.NewWarehouseRepository(s.db, s.logger)
	whUsecase := usecase.NewIWarehouseUsecase(whRepository)
	auUsecase := usecase.NewAuthUsecase(manager)
	auMiddlware := custom_middleware.NewAuthHttpMiddleware(auUsecase, s.cfg)
	whHttpHandler := handler.NewIWareHouseHandler(s.logger, whUsecase)

	zoneRepo := repositories.NewZoneRepository(s.db, s.logger)
	zoneUsecase := usecase.NewZoneUsecase(zoneRepo)
	zoneHttpHandler := handler.NewIZoneHandler(s.logger, zoneUsecase)

	//User routes
	userRouters := group.Group("/auth")
	userRouters.POST("/sign-up", userHttpHandler.Register)
	userRouters.POST("/sign-in-phone", userHttpHandler.LoginByPhoneNumber)
	userRouters.POST("/sign-in-email", userHttpHandler.LoginByEmail)
	userRouters.POST("/refresh", userHttpHandler.Refresh)

	// Wh routes
	whRouters := group.Group("/warehouse")
	whRouters.Use(auMiddlware.Auth, roleMiddleware.IsAdmin)
	whRouters.POST("", whHttpHandler.CreateWarehouse)
	whRouters.GET("", whHttpHandler.GetAllWarehouses)
	whRouters.GET("/:name", whHttpHandler.GetWarehouse)
	whRouters.PUT("/:name", whHttpHandler.UpdateWarehouse)
	whRouters.DELETE("/:name", whHttpHandler.DeleteWarehouse)

	// Zone routes
	zoneRouters := whRouters.Group("/:warehouse_id/zone")
	zoneRouters.Use(auMiddlware.Auth, roleMiddleware.IsAdmin)
	zoneRouters.POST("", zoneHttpHandler.CreateZone)
	zoneRouters.PUT("/:zone_id", zoneHttpHandler.UpdateZone)
	zoneRouters.GET("", zoneHttpHandler.GetAllZones)
	zoneRouters.GET("/:zone_id", zoneHttpHandler.GetZone)
	zoneRouters.DELETE("/:zone_id", zoneHttpHandler.DeleteZone)
}
